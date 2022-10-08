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

type ResBody struct {
	Token string `json:"token"`
}

func getRandInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(999999)
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

	res := ResBody{}
	_ = json.Unmarshal(w.Body.Bytes(), &res)

	return res.Token
}

func TestSignUpShouldHaveStatusCode201(t *testing.T) {
	dbConn := db.Connect("webcuss_test")
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
	dbConn := db.Connect("webcuss_test")
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

	res := ResBody{}
	_ = json.Unmarshal(w.Body.Bytes(), &res)
	assert.NotEmpty(t, res.Token)
}

func TestSignInShouldHaveStatusCode401(t *testing.T) {
	dbConn := db.Connect("webcuss_test")
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
