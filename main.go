// This code taken from
// https://github.com/gin-gonic/gin#quick-start

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRouter() (router *gin.Engine) {
	router = gin.Default()
	return
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	router := getRouter()
	router.GET("/ping", pingHandler)
	router.Run()
}
