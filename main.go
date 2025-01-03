package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	//define gin router
	router := gin.Default()

	//group the routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
		v1.POST("/users", CreateUser)
		v1.PUT("/users/:id", UpdateUser)
		v1.DELETE("/users/:id", DeleteUser)

		v1.GET("/posts", GetPosts)
		v1.GET("/posts/:id", GetPost)
		v1.POST("/posts", CreatePost)
		v1.PUT("/posts/:id", UpdatePost)
		v1.DELETE("/posts/:id", DeletePost)
	}
	//	listen to port
	router.Run(":9000")
}
