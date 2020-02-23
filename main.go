package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niwek/niwek-swagger/controller"
	"github.com/niwek/niwek-swagger/env"
)

func main() {
	router := gin.Default()

	// Ping function
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/v1")
	{
		v1.POST("/user", controller.CreateUser)
		v1.GET("/user/:id", controller.GetUserByID)
	}

	router.Run(":" + env.EnvConfig.Port)
}
