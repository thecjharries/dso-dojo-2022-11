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
