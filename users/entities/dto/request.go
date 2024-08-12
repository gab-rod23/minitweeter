package dto

type CreateUserRequestDTO struct {
	Username *string `json:"username"`
	Name     *string `json:"name"`
	Mail     *string `json:"mail"`
}

type FollowUserRequestDTO struct {
	UsernameToFollow *string `json:"username_to_follow"`
}
