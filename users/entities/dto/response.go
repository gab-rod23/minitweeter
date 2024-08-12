package dto

import "time"

type UserFollowDataResponseDTO struct {
	Following []string
	Followed  []string
}

type UseDataResponseDTO struct {
	Username    string    `json:"username"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	CreatedDate time.Time `json:"created_date"`
	Followers   []string  `json:"followers"`
	Following   []string  `json:"following"`
}
