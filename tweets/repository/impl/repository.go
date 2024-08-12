package impl

import (
	"context"

	"github.com/gab-rod23/minitweeter/database/mongodb"
	"github.com/gab-rod23/minitweeter/tweets/entities/dto"
	"github.com/gab-rod23/minitweeter/tweets/entities/model"
	"github.com/gab-rod23/minitweeter/tweets/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (t tweetRepository) FindTweetsFromUsers(timelineData *dto.TimelineTweetData, followingUsers []string) ([]model.TweetModelCollection, error) {
	tweetCollection := t.client.GetCollection(TWEET_COLLECTION_NAME)

	var limitValue int64
	var skipValue int64
	limitValue = int64(timelineData.PageSize)
	skipValue = int64(timelineData.PageSize * timelineData.PageNumber)
	pageOptions := options.FindOptions{Limit: &limitValue, Skip: &skipValue}
	sortOptions := options.Find()
	sortOptions.SetSort(bson.D{{"created_date", -1}})

	conditions := bson.D{}
	for _, user := range followingUsers {
		conditions = append(conditions, bson.E{Key: "username", Value: user})
	}
	cur, err := tweetCollection.Find(context.TODO(), conditions, &pageOptions, sortOptions)
	if err != nil {
		return nil, err
	}
	result := new([]model.TweetModelCollection)
	err = cur.All(context.TODO(), result)
	if err != nil {
		return nil, err
	}
	return *result, nil
}
