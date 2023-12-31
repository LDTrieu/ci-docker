package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"code": 200})
	})

	router.Run(":8080")
}
