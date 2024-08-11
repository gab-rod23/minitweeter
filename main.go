package main

import (
	"net/http"

	"github.com/gab-rod23/minitweeter/database/mongodb"
	userControllerImpl "github.com/gab-rod23/minitweeter/users/controller/impl"
	"github.com/gin-gonic/gin"
)

func main() {
	err := mongodb.InitConnection()
	if err != nil {
		panic(err)
	}
	userController := userControllerImpl.NewUserController()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	userRouterGroup := r.Group("/user")
	userRouterGroup.POST("/create", userController.HandlerCreateNewUser)
	userRouterGroup.GET("", userController.HandlerRetrieveUserData)

	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
