package impl

import (
	"context"

	"github.com/gab-rod23/minitweeter/database/mongodb"
	"github.com/gab-rod23/minitweeter/users/entities/model"
	"github.com/gab-rod23/minitweeter/users/repository"
	"github.com/gab-rod23/minitweeter/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const USERS_COLLECTION_NAME = "users"

type userRepository struct {
	client *mongodb.MongoDBConnection
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{
		client: mongodb.GetClient(),
	}
}

func (r userRepository) InsertUser(newUser *model.UserModelCollection) error {
	userCollection := r.client.GetCollection(USERS_COLLECTION_NAME)
	_, err := userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return err
	}
	return nil
}

func (r userRepository) FindUserByField(value string, field string) (*model.UserModelCollection, error) {
	userData := new(model.UserModelCollection)
	userCollection := r.client.GetCollection(USERS_COLLECTION_NAME)
	err := userCollection.FindOne(context.TODO(), bson.D{{field, value}}).Decode(userData)
	if err != nil {
		if mongo.ErrNoDocuments == err {
			return nil, util.ErrUserNotFound
		}
		return nil, err
	}
	return userData, nil
}

func (r userRepository) AddtNewFollowerToUser(valueFilter string, fieldFilter string, followerUsername string) error {
	return r.pushValueIntoArray(valueFilter, fieldFilter, "followers", followerUsername)
}

func (r userRepository) AddNewFollowingToUser(valueFilter string, fieldFilter string, followingUsername string) error {
	return r.pushValueIntoArray(valueFilter, fieldFilter, "following", followingUsername)
}

func (r userRepository) pushValueIntoArray(valueFilter string, fieldFilter string, arrayName string, valueToPush string) error {
	userCollection := r.client.GetCollection(USERS_COLLECTION_NAME)
	_, err := userCollection.UpdateOne(context.TODO(), bson.D{{Key: fieldFilter, Value: valueFilter}}, bson.D{{"$push", bson.D{{arrayName, valueToPush}}}})
	return err
}
