package repository

import (
	"github.com/gab-rod23/minitweeter/tweets/entities/dto"
	"github.com/gab-rod23/minitweeter/tweets/entities/model"
)

type TweetRepository interface {
	InsertTweet(newTweet *model.TweetModelCollection) error
	FindTweetsFromUsers(timelineData *dto.TimelineTweetData, followingUsers []string) ([]model.TweetModelCollection, error)
}
