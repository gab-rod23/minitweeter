package impl

import (
	"context"
	"fmt"

	"github.com/gab-rod23/minitweeter/database/mongodb"
	"github.com/gab-rod23/minitweeter/tweets/entities/dto"
	"github.com/gab-rod23/minitweeter/tweets/entities/model"
	"github.com/gab-rod23/minitweeter/tweets/repository"
	"go.mongodb.org/mongo-driver/bson"

	//	"go.mongodb.org/mongo-driver/bson/primitive"
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

	usernameList := bson.A{}
	for _, user := range followingUsers {
		usernameList = append(usernameList, bson.D{{"username", user}})
	}
	fmt.Println(usernameList)
	filterUsernames := bson.D{{"$or", usernameList}}
	//var filters primitive.D
	// if timelineData.LastPageTweetDate != nil {
	// 	lowerDateFilter := bson.D{{"created_date", bson.D{{"$lt", timelineData.LastPageTweetDate}}}}
	// 	filters = bson.D{{"$and", bson.A{lowerDateFilter, filterUsernames}}}
	// } else {
	// 	filters = filterUsernames
	// }

	//cur, err := tweetCollection.Find(context.TODO(), filters, &pageOptions, sortOptions)
	fmt.Println(pageOptions)
	cur, err := tweetCollection.Find(context.TODO(), filterUsernames, &pageOptions, sortOptions)
	if err != nil {
		return nil, err
	}
	result := new([]model.TweetModelCollection)
	fmt.Println(cur)
	err = cur.All(context.TODO(), result)
	if err != nil {
		return nil, err
	}
	fmt.Print(result)
	return *result, nil
}
