package dto

import "time"

type CreateTweetRequestDto struct {
	Text string `json:"text"`
}

type TimelineTweetData struct {
	Username          string
	PageSize          int
	PageNumber        int
	LastPageTweetDate *time.Time
}
