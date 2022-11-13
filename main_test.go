package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRouter(t *testing.T) {
	router := getRouter()
	assert.NotNil(t, router, "The router should not be nil")
}

func TestPingHandler(t *testing.T) {
	router := getRouter()
	router.GET("/ping", pingHandler)
	request := httptest.NewRequest("GET", "/ping", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "The response code should be 200")
	var body map[string]string
	json.Unmarshal(response.Body.Bytes(), &body)
	assert.Equal(t, "pong", body["message"], "The response body should be 'pong'")
}

func TestVersionHandler(t *testing.T) {
	router := getRouter()
	router.GET("/version", versionHandler)
	request := httptest.NewRequest("GET", "/version", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "The response code should be 200")
	var body map[string]string
	json.Unmarshal(response.Body.Bytes(), &body)
	assert.Equal(t, "-1", body["version"], "The response body should be '-1'")
}

func TestMain(t *testing.T) {
	go main()
	response, err := http.Get("http://localhost:8080/ping")
	assert.Nil(t, err, "The response should not be nil")
	assert.Equal(t, http.StatusOK, response.StatusCode, "The response code should be 200")
	var body map[string]string
	json.NewDecoder(response.Body).Decode(&body)
	assert.Equal(t, "pong", body["message"], "The response body should be 'pong'")
}
