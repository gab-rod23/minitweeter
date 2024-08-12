package impl

import (
	"context"

	"github.com/gab-rod23/minitweeter/database/mongodb"
	"github.com/gab-rod23/minitweeter/tweets/entities/model"
	"github.com/gab-rod23/minitweeter/tweets/repository"
)

const TWEET_COLLECTION_NAME = "tweets"

type tweetRepository struct {
	client *mongodb.MongoDBConnection
}

func NewTweetRepository() repository.TweetRepository {
	return &tweetRepository{
		client: mongodb.GetClient(),
	}
}

func (t tweetRepository) InsertTweet(newTweet *model.TweetModelCollection) error {
	tweetCollection := t.client.GetCollection(TWEET_COLLECTION_NAME)
	_, err := tweetCollection.InsertOne(context.TODO(), newTweet)
	if err != nil {
		return err
	}
	return nil
}
