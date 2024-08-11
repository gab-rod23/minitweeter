package impl

import (
	"context"
	"fmt"

	"github.com/gab-rod23/minitweeter/database/mongodb"
	"github.com/gab-rod23/minitweeter/users/entities/model"
	"github.com/gab-rod23/minitweeter/users/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type userRepository struct {
	client *mongodb.MongoDBConnection
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{
		client: mongodb.GetClient(),
	}
}

func (r userRepository) InsertUser(newUser *model.UserModelCollection) {
	userCollection := r.client.GetCollection("users")
	fmt.Println(newUser.CreatedDate)
	_, err := userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (r userRepository) FindUserByField(value string, field string) (*model.UserModelCollection, error) {
	userData := new(model.UserModelCollection)
	userCollection := r.client.GetCollection("users")
	err := userCollection.FindOne(context.TODO(), bson.D{{field, value}}).Decode(userData)
	if err != nil {
		fmt.Println(err.Error())
	}
	return userData, nil
}
