package usecase

import (
	"github.com/gab-rod23/minitweeter/tweets/entities/dto"
)

type TweetUsecase interface {
	CreateNewTweet(tweet *dto.CreateTweetRequestDto, username string) error
	RetrieveTimelineTweet(*dto.TimelineTweetData) (*dto.TimelineTweetResponseDto, error)
}
