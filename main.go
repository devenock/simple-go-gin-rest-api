package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	//define gin router
	router := gin.Default()

	//	sample route
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello Golang!"})
	})

	//	listen to port
	router.Run(":9000")
}
