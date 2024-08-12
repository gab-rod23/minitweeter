package repository

import "github.com/gab-rod23/minitweeter/tweets/entities/model"

type TweetRepository interface {
	InsertTweet(newTweet *model.TweetModelCollection) error
}
