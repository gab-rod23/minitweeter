package dto

import "time"

type TimelineTweetResponseItem struct {
	Text        string    `json:"text"`
	Username    string    `json:"username"`
	CreatedDate time.Time `json:"created_date"`
}

type TimelineTweetResponseDto struct {
	Timeline []TimelineTweetResponseItem
}
