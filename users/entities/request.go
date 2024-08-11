package entities

type CreateUserRequestDTO struct {
	Username *string `json:"username" binding:"required"`
	Name     *string `json:"name" binding:"required"`
	Mail     *string `json:"mail" binding:"email"`
}

type FollowUserRequestDTO struct {
	UsernameToFollow *string `json:"username_to_follow" binding:"required"`
}
