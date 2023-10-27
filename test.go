package main

import (
	"bytes"
	"json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	// Testing
	"github.com/stretchr/testify/assert"
)

type Request struct {
	City1  string `json:"city1"`
	City2  string `json:"city2"`
	Method string `json:"method"`
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/distance", ginService)
	return router
}

func getRequest(request *Request) *http.Request {
	json_req, json_er := json.Marshal(*request)
	if json_er != nil {
		log.Fatal(json_er)
	}
	req, req_err := http.NewRequest("POST", "/distance", bytes.NewBuffer(json_req))
	if req_err != nil {
		log.Fatal(req_err)
	}
	return req
}

func performRequest(r http.Handler, req *http.Request) (int, string) {
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	return recorder.Code, recorder.Body.String()
}

func TestParams(t *testing.T) {
	// Required params
	// city1 := gin_ctx.Query("city1")
	// city2 := gin_ctx.Query("city2")
	//
	// method := gin_ctx.Query("method")

	router := setupRouter()

	test_req := Request{
		City1:  "London",
		City2:  "Paris",
		Method: "CSV",
	}

	req := getRequest(&test_req)

	code, body := performRequest(router, req)
}
