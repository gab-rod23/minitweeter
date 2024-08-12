package usecase

import (
	"github.com/gab-rod23/minitweeter/tweets/entities/dto"
	"github.com/gin-gonic/gin"
)

type TweetUsecase interface {
	CreateNewTweet(tweet *dto.CreateTweetRequestDto, username string) error
	RetrieveTimelineTweet(ctx *gin.Context) (*dto.TimelineTweetResponseDto, error)
}
