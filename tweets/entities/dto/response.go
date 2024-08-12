package dto

type TimelineTweetResponseItem struct {
	Text        string `json:"text"`
	Username    string `json:"username"`
	CreatedDate string `json:"created_date"`
}

type TimelineTweetResponseDto struct {
	timeline []TimelineTweetResponseItem
}
