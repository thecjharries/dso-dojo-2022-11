package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRouter(t *testing.T) {
	router := getRouter()
	assert.NotNil(t, router, "The router should not be nil")
}
