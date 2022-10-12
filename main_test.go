package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/webcuss/webcuss/db"
	"github.com/webcuss/webcuss/route"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const testDbName = "webcuss"

type ResAuthBody struct {
	Token string `json:"token"`
}

type ResPostTopicBody struct {
	Id        string `field:"id"`
	CommentId string `field:"commentId"`
}

type ResGetTopicBody struct {
	Data []any `field:"data"`
	Pg   int   `field:"pg"`
}

type ResPostCommentBody struct {
	Id string `field:"id"`
}

type ResUser struct {
	Id    string `field:"id"`
	Uname string `field:"uname"`
}

type ResTopicComment struct {
	Id      string  `field:"id"`
	Content string  `field:"content"`
	User    ResUser `field:"user"`
}

type ResGetTopicComments struct {
	Pg   int               `field:"pg"`
	Data []ResTopicComment `field:"data"`
}

func getRandInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(999999)
}

func getRandomString() string {
	return fmt.Sprintf("%d", getRandInt())
}

func signUp(t *testing.T, r *gin.Engine, uname, pword string) string {
	body := gin.H{
		"uname": uname,
		"pword": pword,
	}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sup", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	res := ResAuthBody{}
	_ = json.Unmarshal(w.Body.Bytes(), &res)

	return res.Token
}

func TestSignUpShouldHaveStatusCode201(t *testing.T) {
	dbConn := db.Connect()
	defer dbConn.Close()
	db.CreateTables(dbConn)

	router := route.SetupRouter(dbConn)

	var (
		uname = fmt.Sprintf("john%d", getRandInt())
		pword = "123456"
	)
	token := signUp(t, router, uname, pword)
	assert.NotEmpty(t, token)
}

func TestSignInShouldHaveStatusCode200(t *testing.T) {
	dbConn := db.Connect()
	defer dbConn.Close()
	db.CreateTables(dbConn)

	router := route.SetupRouter(dbConn)

	rand.Seed(time.Now().UnixNano())
	var (
		uname = fmt.Sprintf("john%d", getRandInt())
		pword = "123456"
	)
	log.Println("uname=", uname)
	_ = signUp(t, router, uname, pword)

	body := gin.H{
		"uname": uname,
		"pword": pword,
	}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sin", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	res := ResAuthBody{}
	_ = json.Unmarshal(w.Body.Bytes(), &res)
	assert.NotEmpty(t, res.Token)
}

func TestSignInShouldHaveStatusCode401(t *testing.T) {
	dbConn := db.Connect()
	defer dbConn.Close()
	db.CreateTables(dbConn)

	router := route.SetupRouter(dbConn)

	body := gin.H{
		"uname": "nonexistentuser898790707790",
		"pword": "123456xxxxxxxxx",
	}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sin", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "Incorrect credentials", w.Body.String())
}

func TestPostTopicShouldHaveExpectedResultsWhenHasComment(t *testing.T) {
	dbConn := db.Connect()
	defer dbConn.Close()
	db.CreateTables(dbConn)

	router := route.SetupRouter(dbConn)

	randHostname := "https://" + getRandomString() + ".example.com/category/blah/page.php?p1=abc&p2=123"
	body := gin.H{
		"url":     randHostname,
		"title":   "Lorem ipsum",
		"comment": "first comment!",
	}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tpc", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	randomUser := fmt.Sprintf("user%d", getRandInt())
	req.Header.Set("Authorization", "Bearer "+signUp(t, router, randomUser, "123456"))
	router.ServeHTTP(w, req)

	log.Println(w.Body.String())
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, w.Body.String())

	res := ResPostTopicBody{}
	_ = json.Unmarshal(w.Body.Bytes(), &res)
	assert.NotEmpty(t, res.Id)
	assert.NotEmpty(t, res.CommentId)
}

func TestPostTopicShouldHaveExpectedResultsWhenNoComment(t *testing.T) {
	dbConn := db.Connect()
	defer dbConn.Close()
	db.CreateTables(dbConn)

	router := route.SetupRouter(dbConn)

	randHostname := "https://" + getRandomString() + ".example.com/category/blah/page.php?p1=abc&p2=123"
	body := gin.H{
		"url":   randHostname,
		"title": "Lorem ipsum",
	}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tpc", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	randomUser := fmt.Sprintf("user%d", getRandInt())
	req.Header.Set("Authorization", "Bearer "+signUp(t, router, randomUser, "123456"))
	router.ServeHTTP(w, req)

	log.Println(w.Body.String())
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, w.Body.String())

	res := ResPostTopicBody{}
	_ = json.Unmarshal(w.Body.Bytes(), &res)
	assert.NotEmpty(t, res.Id)
	assert.Empty(t, res.CommentId)
}

func TestPostTopicShouldReturnSameIdWhenUrlIsDuplicate(t *testing.T) {
	dbConn := db.Connect()
	defer dbConn.Close()
	db.CreateTables(dbConn)

	router := route.SetupRouter(dbConn)

	randHostname := "https://" + getRandomString() + ".example.com/category/blah/page.php?p1=abc&p2=123"
	body := gin.H{
		"url":     randHostname,
		"title":   "Lorem ipsum",
		"comment": "comment" + getRandomString(),
	}
	b, _ := json.Marshal(body)

	randomUser := fmt.Sprintf("user%d", getRandInt())
	authToken := "Bearer " + signUp(t, router, randomUser, "123456")

	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/tpc", bytes.NewReader(b))
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("Authorization", authToken)
	router.ServeHTTP(w1, req1)

	log.Println("w1", w1.Body.String())
	assert.Equal(t, http.StatusCreated, w1.Code)

	res1 := ResPostTopicBody{}
	_ = json.Unmarshal(w1.Body.Bytes(), &res1)

	tpcId := res1.Id
	assert.NotEmpty(t, tpcId)

	// request again
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("POST", "/tpc", bytes.NewReader(b))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("Authorization", authToken)
	router.ServeHTTP(w2, req2)
	log.Println("w2", w2.Body.String())
	assert.Equal(t, http.StatusCreated, w2.Code)

	res2 := ResPostTopicBody{}
	_ = json.Unmarshal(w2.Body.Bytes(), &res2)

	assert.Equal(t, tpcId, res2.Id)
}

func TestGetTopicShouldHaveExpectedResult(t *testing.T) {
	dbConn := db.Connect()
	defer dbConn.Close()
	db.CreateTables(dbConn)

	router := route.SetupRouter(dbConn)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tpc", nil)
	req.Header.Set("Content-Type", "application/json")
	randomUser := fmt.Sprintf("user%d", getRandInt())
	req.Header.Set("Authorization", "Bearer "+signUp(t, router, randomUser, "123456"))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())

	var res ResGetTopicBody
	_ = json.Unmarshal(w.Body.Bytes(), &res)
	assert.Greater(t, res.Pg, 0)
	assert.NotNil(t, res.Data)
}

func TestPostCommentShouldHaveExpectedResult(t *testing.T) {
	dbConn := db.Connect()
	defer dbConn.Close()
	db.CreateTables(dbConn)

	router := route.SetupRouter(dbConn)

	authToken := signUp(t, router, fmt.Sprintf("user%d", getRandInt()), "123456")

	randHostname := "https://" + getRandomString() + ".example.com/category/blah/page.php?p1=abc&p2=123"
	body1 := gin.H{
		"url":   randHostname,
		"title": "Lorem ipsum",
	}
	b1, _ := json.Marshal(body1)

	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/tpc", bytes.NewReader(b1))
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("Authorization", "Bearer "+authToken)
	router.ServeHTTP(w1, req1)

	assert.Equal(t, http.StatusCreated, w1.Code)
	assert.NotEmpty(t, w1.Body.String())

	res1 := ResPostTopicBody{}
	_ = json.Unmarshal(w1.Body.Bytes(), &res1)
	assert.NotEmpty(t, res1.Id)

	// post comment
	body2 := gin.H{
		"comment": "I like it! " + getRandomString(),
	}
	b2, _ := json.Marshal(body2)

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("POST", fmt.Sprintf("/tpc/%s/cmt", res1.Id), bytes.NewReader(b2))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("Authorization", "Bearer "+authToken)
	router.ServeHTTP(w2, req2)

	assert.Equal(t, http.StatusCreated, w2.Code)
	var res2 ResPostCommentBody
	_ = json.Unmarshal(w2.Body.Bytes(), &res2)
	assert.NotEmpty(t, res2.Id)
}

func TestPostCommentShouldHave401BadRequestWhenTopicIdIsInvalid(t *testing.T) {
	dbConn := db.Connect()
	defer dbConn.Close()
	db.CreateTables(dbConn)

	router := route.SetupRouter(dbConn)

	authToken := signUp(t, router, fmt.Sprintf("user%d", getRandInt()), "123456")

	body := gin.H{
		"comment": "I like it! " + getRandomString(),
	}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	invalidTopicId := "xxxx" + getRandomString()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/tpc/%s/cmt", invalidTopicId), bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, w.Body.String(), "Invalid topicId")
}

func TestPostCommentShouldHave404NotFoundWhenTopicIdIsNonExistent(t *testing.T) {
	dbConn := db.Connect()
	defer dbConn.Close()
	db.CreateTables(dbConn)

	router := route.SetupRouter(dbConn)

	authToken := signUp(t, router, fmt.Sprintf("user%d", getRandInt()), "123456")

	body := gin.H{
		"comment": "I like it! " + getRandomString(),
	}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	invalidTopicId := "d299f62e-99ce-4000-acb2-55184d1a835f"
	req, _ := http.NewRequest("POST", fmt.Sprintf("/tpc/%s/cmt", invalidTopicId), bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetCommentShouldHaveExpectedResult(t *testing.T) {
	dbConn := db.Connect()
	defer dbConn.Close()
	db.CreateTables(dbConn)

	router := route.SetupRouter(dbConn)

	uname := fmt.Sprintf("user%d", getRandInt())
	authToken := signUp(t, router, uname, "123456")

	// create topic
	randHostname := "https://" + getRandomString() + ".example.com/category/blah/page.php?p1=abc&p2=123"
	bodyTopic := gin.H{
		"url":   randHostname,
		"title": "Lorem ipsum",
	}
	bTopic, _ := json.Marshal(bodyTopic)

	wTopic := httptest.NewRecorder()
	reqTopic, _ := http.NewRequest("POST", "/tpc", bytes.NewReader(bTopic))
	reqTopic.Header.Set("Content-Type", "application/json")
	reqTopic.Header.Set("Authorization", "Bearer "+authToken)
	router.ServeHTTP(wTopic, reqTopic)

	var resTopic ResPostTopicBody
	_ = json.Unmarshal(wTopic.Body.Bytes(), &resTopic)
	topicId := resTopic.Id

	assert.Equal(t, http.StatusCreated, wTopic.Code)
	assert.NotEmpty(t, topicId)

	// post comments

	cmtIds := make([]string, 0)
	loopLen := 10
	for i := 0; i < loopLen; i++ {
		bodyComment := gin.H{
			"comment": "I like it! " + getRandomString(),
		}
		bComment, _ := json.Marshal(bodyComment)

		wComment := httptest.NewRecorder()
		reqComment, _ := http.NewRequest("POST", fmt.Sprintf("/tpc/%s/cmt", topicId), bytes.NewReader(bComment))
		reqComment.Header.Set("Content-Type", "application/json")
		reqComment.Header.Set("Authorization", "Bearer "+authToken)
		router.ServeHTTP(wComment, reqComment)

		assert.Equal(t, http.StatusCreated, wComment.Code)

		var resComment ResPostCommentBody
		_ = json.Unmarshal(wComment.Body.Bytes(), &resComment)
		assert.NotEmpty(t, resComment.Id)

		cmtIds = append(cmtIds, resComment.Id)
	}

	assert.Equal(t, loopLen, len(cmtIds))

	wTopicComments := httptest.NewRecorder()
	reqTopicComments, _ := http.NewRequest("GET", fmt.Sprintf("/tpc/%s/cmt", topicId), nil)
	reqTopicComments.Header.Set("Content-Type", "application/json")
	reqTopicComments.Header.Set("Authorization", "Bearer "+authToken)
	router.ServeHTTP(wTopicComments, reqTopicComments)

	assert.Equal(t, http.StatusOK, wTopicComments.Code)

	resTopicComments := ResGetTopicComments{}
	_ = json.Unmarshal(wTopicComments.Body.Bytes(), &resTopicComments)
	assert.Equal(t, loopLen, len(resTopicComments.Data))

	// assert comments are posted
	for _, cmt := range resTopicComments.Data {
		// assert user
		assert.Equal(t, uname, cmt.User.Uname)

		exists := false
		for _, cmtId := range cmtIds {
			if cmtId == cmt.Id {
				exists = true
				break
			}
		}
		if !exists {
			assert.Fail(t, "comment id not exists")
			break
		}
	}
}

func TestGetCommentShouldHave404NotFoundWhenTopicIdIsNonExistent(t *testing.T) {
	dbConn := db.Connect()
	defer dbConn.Close()
	db.CreateTables(dbConn)

	router := route.SetupRouter(dbConn)

	uname := fmt.Sprintf("user%d", getRandInt())
	authToken := signUp(t, router, uname, "123456")
	w := httptest.NewRecorder()
	nonExistentTopicId := "9e665c2a-cc4c-4844-9788-7a42495f0acb"
	req, _ := http.NewRequest("GET", fmt.Sprintf("/tpc/%s/cmt", nonExistentTopicId), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
