package usecase

import (
	"github.com/gab-rod23/minitweeter/users/entities/dto"
)

type UserUsecase interface {
	CreateNewUser(newUserData *dto.CreateUserRequestDTO) error
	FollowUser(username string, followUserData *dto.FollowUserRequestDTO) error
	RetrieveUserByUsername(username string) (*dto.UseDataResponseDTO, error)
}
