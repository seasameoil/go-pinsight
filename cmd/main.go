package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/seasameoil/go-pinsight.git/model"
)

func main() {
	model.InitDB()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
