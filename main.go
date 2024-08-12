package main

import (
	"net/http"

	"github.com/gab-rod23/minitweeter/database/mongodb"
	tweetControllerImpl "github.com/gab-rod23/minitweeter/tweets/controller/impl"
	userControllerImpl "github.com/gab-rod23/minitweeter/users/controller/impl"
	"github.com/gin-gonic/gin"
)

func main() {
	err := mongodb.InitConnection()
	if err != nil {
		panic(err)
	}
	userController := userControllerImpl.NewUserController()
	tweetController := tweetControllerImpl.NewTweetController()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	userRouterGroup := r.Group("/user")
	{
		userRouterGroup.POST("/create", userController.HandlerCreateNewUser)
		userRouterGroup.GET("/read", userController.HandlerRetrieveUserDataByUsername)
		userRouterGroup.POST("/follow", userController.HandlerFollowUser)
	}

	tweetRouterGroup := r.Group("/tweet")
	{
		tweetRouterGroup.POST("/create", tweetController.HandlerCreateNewTweet)
		tweetRouterGroup.GET("/timeline", tweetController.HandlerRetrieveTimelineTweet)
	}

	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
