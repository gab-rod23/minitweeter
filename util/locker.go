package util

import (
	"context"
	"errors"
	"time"

	"github.com/gab-rod23/minitweeter/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const LOCK_RETRIES = 20

func Lock(field string, value string, collectionName string) error {
	collection := mongodb.GetClient().GetCollection("lock")
	retries := 0
	_, err := collection.InsertOne(context.TODO(), bson.D{{field, value}, {"collection_name", collectionName}})
	for err != nil && mongo.IsDuplicateKeyError(err) {
		if retries > LOCK_RETRIES {
			return errors.New("Lock timeout")
		}
		time.Sleep(500 * time.Microsecond)
		_, err = collection.InsertOne(context.TODO(), bson.D{{field, value}, {"collection_name", collectionName}})
		retries++
	}
	return nil
}

func Unlock(field string, value string, collectionName string) {
	collection := mongodb.GetClient().GetCollection("lock")
	collection.DeleteOne(context.TODO(), bson.D{{field, value}, {"collection_name", collectionName}})
}
