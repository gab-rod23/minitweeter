package impl

import (
	"context"
	"errors"
	"fmt"
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
	u.userRepository.InsertUser(userToInsert)
	return nil
}

func (u userUsecase) FollowUser(username string, followUserData *dto.FollowUserRequestDTO) error {

	usernameToFollow := *followUserData.UsernameToFollow
	if errLock := util.Lock(USERNAME_FIELD, usernameToFollow, impl.USERS_COLLECTION_NAME); errLock != nil {
		return errLock
	}
	defer util.Unlock(USERNAME_FIELD, usernameToFollow, impl.USERS_COLLECTION_NAME)
	userDataToFollow, err := u.userRepository.FindUserByField(usernameToFollow, USERNAME_FIELD)
	if userDataToFollow == nil {
		return errors.New("usuario a seguir inexistente")
	}
	if err != nil {
		return errors.New("error al recuperar el usuario a seguir")
	}

	if errLock := util.Lock(USERNAME_FIELD, username, impl.USERS_COLLECTION_NAME); errLock != nil {
		return errLock
	}
	defer util.Unlock(USERNAME_FIELD, username, impl.USERS_COLLECTION_NAME)
	userData, _ := u.userRepository.FindUserByField(username, USERNAME_FIELD)
	for _, userFollowing := range userData.Following {
		fmt.Print("Usuario siguiendo ")
		fmt.Println(userFollowing)
		if usernameToFollow == userFollowing {
			return nil
		}
	}
	session, err := mongodb.StartTransaction(context.TODO())
	if err != nil {
		return err
	}
	userData.Following = append(userData.Following, usernameToFollow)
	fmt.Println(userData.Following)
	userDataToFollow.Followers = append(userDataToFollow.Followers, username)
	fmt.Println(userDataToFollow.Followers)
	if err := u.userRepository.AddNewFollowingToUser(username, USERNAME_FIELD, usernameToFollow); err != nil {
		fmt.Println("error en following")
		mongodb.RollbackTransaction(context.Background(), session)
		return err
	}
	if err := u.userRepository.AddtNewFollowerToUser(usernameToFollow, USERNAME_FIELD, username); err != nil {
		fmt.Println("error en followers")
		mongodb.RollbackTransaction(context.Background(), session)
		return err
	}
	fmt.Println("termino ok")
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
