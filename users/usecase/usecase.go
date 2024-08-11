package usecase

import "github.com/gab-rod23/minitweeter/users/entities"

type UserUsecase interface {
	CreateNewUser(newUserData *entities.CreateUserRequestDTO) error
	FollowUser(followUserData *entities.FollowUserRequestDTO) error
	RetrieveUserByUsername(username string) (*entities.UseDataResponseDTO, error)
}
