package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TweetModelCollection struct {
	ID          primitive.ObjectID `bson:"_id, omitempty" json:"id,omitempty"`
	Username    string             `bson:"username" json:"username"`
	Text        string             `bson:"text" json:"text"`
	CreatedDate time.Time          `bson:"created_date" json:"created_date"`
}
