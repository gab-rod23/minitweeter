package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTweetRequestDto struct {
	Text string `json:"text"`
}

type TimelineTweetData struct {
	Username          string
	PageSize          int
	PageNumber        int
	LastPageTweetDate *primitive.DateTime
}
