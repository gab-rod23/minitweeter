package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModelCollection struct {
	ID          primitive.ObjectID `bson:"_id, omitempty" json:"id,omitempty"`
	Username    string             `bson:"username" json:"username"`
	Name        string             `bson:"name" json:"name"`
	Email       string             `bson:"email" json:"email"`
	CreatedDate time.Time          `bson:"created_date" json:"created_date"`
	Followers   []string           `json:"followers"`
	Following   []string           `json:"following"`
}
