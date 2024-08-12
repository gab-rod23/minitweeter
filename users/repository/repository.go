package repository

import (
	"github.com/gab-rod23/minitweeter/users/entities/model"
)

type UserRepository interface {
	InsertUser(newUser *model.UserModelCollection) error
	FindUserByField(value string, field string) (*model.UserModelCollection, error)
	AddtNewFollowerToUser(valueFilter string, fieldFilter string, followerUsername string) error
	AddNewFollowingToUser(valueFilter string, fieldFilter string, followingUsername string) error
}
