package impl

import (
	"time"

	"github.com/gab-rod23/minitweeter/users/entities"
	"github.com/gab-rod23/minitweeter/users/entities/model"
	"github.com/gab-rod23/minitweeter/users/repository"
	"github.com/gab-rod23/minitweeter/users/repository/impl"
	"github.com/gab-rod23/minitweeter/users/usecase"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	USERNAME_FIELD = "username"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase() usecase.UserUsecase {
	return &userUsecase{
		userRepository: impl.NewUserRepository(),
	}
}

func (u userUsecase) CreateNewUser(newUserData *entities.CreateUserRequestDTO) error {
	userToInsert := generateUser(newUserData)
	u.userRepository.InsertUser(userToInsert)
	return nil
}

func (u userUsecase) FollowUser(followUserData *entities.FollowUserRequestDTO) error {
	return nil
}

func (u userUsecase) RetrieveUserByUsername(username string) (*entities.UseDataResponseDTO, error) {
	userData, _ := u.userRepository.FindUserByField(username, USERNAME_FIELD)
	userResponse := &entities.UseDataResponseDTO{
		Username:    userData.Username,
		Name:        userData.Name,
		Email:       userData.Email,
		CreatedDate: userData.CreatedDate.Local(),
		Followers:   userData.Followers,
		Following:   userData.Following,
	}
	return userResponse, nil
}

func generateUser(userData *entities.CreateUserRequestDTO) *model.UserModelCollection {
	return &model.UserModelCollection{
		ID:          primitive.NewObjectID(),
		Username:    *userData.Username,
		Name:        *userData.Name,
		Email:       *userData.Mail,
		CreatedDate: time.Now(),
		Followers:   []string{},
		Following:   []string{},
	}
}
