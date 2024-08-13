package impl

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/gab-rod23/minitweeter/database/mongodb"
	"github.com/gab-rod23/minitweeter/users/entities/dto"
	"github.com/gab-rod23/minitweeter/users/entities/model"
	"github.com/gab-rod23/minitweeter/users/repository"
	"github.com/gab-rod23/minitweeter/users/repository/impl"
	"github.com/gab-rod23/minitweeter/users/usecase"
	"github.com/gab-rod23/minitweeter/util"
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

func (u userUsecase) CreateNewUser(newUserData *dto.CreateUserRequestDTO) error {
	userToInsert := generateUser(newUserData)
	err := u.userRepository.InsertUser(userToInsert)
	return getDetailedError(err)
}

func (u userUsecase) FollowUser(username string, followUserData *dto.FollowUserRequestDTO) error {

	usernameToFollow := *followUserData.UsernameToFollow
	if errLock := util.Lock(USERNAME_FIELD, usernameToFollow, util.USERS_COLLECTION_NAME); errLock != nil {
		return errLock
	}
	defer util.Unlock(USERNAME_FIELD, usernameToFollow, util.USERS_COLLECTION_NAME)
	userDataToFollow, err := u.userRepository.FindUserByField(usernameToFollow, USERNAME_FIELD)
	if err != nil {
		if errors.Is(err, util.ErrUserNotFound) {
			return util.ErrUserToFollowNotFound
		}
		return err
	}

	if errLock := util.Lock(USERNAME_FIELD, username, util.USERS_COLLECTION_NAME); errLock != nil {
		return errLock
	}
	defer util.Unlock(USERNAME_FIELD, username, util.USERS_COLLECTION_NAME)
	userData, err := u.userRepository.FindUserByField(username, USERNAME_FIELD)
	if err != nil {
		return err
	}
	for _, userFollowing := range userData.Following {
		if usernameToFollow == userFollowing {
			return nil
		}
	}
	session, err := mongodb.StartTransaction(context.TODO())
	if err != nil {
		return err
	}
	userData.Following = append(userData.Following, usernameToFollow)
	userDataToFollow.Followers = append(userDataToFollow.Followers, username)
	if err := u.userRepository.AddNewFollowingToUser(username, USERNAME_FIELD, usernameToFollow); err != nil {
		mongodb.RollbackTransaction(context.Background(), session)
		return err
	}
	if err := u.userRepository.AddtNewFollowerToUser(usernameToFollow, USERNAME_FIELD, username); err != nil {
		mongodb.RollbackTransaction(context.Background(), session)
		return err
	}
	mongodb.CommitTransaction(context.Background(), session)

	return nil
}

func (u userUsecase) RetrieveUserByUsername(username string) (*dto.UseDataResponseDTO, error) {
	userData, _ := u.userRepository.FindUserByField(username, USERNAME_FIELD)
	userResponse := &dto.UseDataResponseDTO{
		Username:    userData.Username,
		Name:        userData.Name,
		Email:       userData.Email,
		CreatedDate: userData.CreatedDate.Local(),
		Followers:   userData.Followers,
		Following:   userData.Following,
	}
	return userResponse, nil
}

func generateUser(userData *dto.CreateUserRequestDTO) *model.UserModelCollection {
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

func getDetailedError(err error) error {
	if err != nil && strings.Contains(err.Error(), "duplicate key error") {
		if strings.Contains(err.Error(), "index: email") {
			return util.ErrEmailAlreadyExists
		}
		if strings.Contains(err.Error(), "index: username") {
			return util.ErrUsernameAlreadyExists
		}
	}
	return err
}
