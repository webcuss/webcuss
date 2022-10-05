package main_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/webcuss/webcuss/db"
	"github.com/webcuss/webcuss/route"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSignupShouldHaveStatusCode201(t *testing.T) {
	dbConn := db.SetupDatabase("webcuss_test")
	defer dbConn.Close()
	db.CreateDatabaseTables(dbConn)

	router := route.SetupRouter(dbConn)

	body := gin.H{
		"uname": "john" + strconv.FormatInt(time.Now().Unix(), 10),
		"pword": "123456",
	}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sup", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	log.Println(w.Body.String())

	type ResBody struct {
		Token string `json:"token"`
	}

	res := ResBody{}
	_ = json.Unmarshal(w.Body.Bytes(), &res)
	assert.NotEmpty(t, res.Token)
}
