package controller

import "github.com/gin-gonic/gin"

type TweetController interface {
	HandlerCreateNewTweet(ctx *gin.Context)
	HandlerRetrieveTimelineTweet(ctx *gin.Context)
}
