package repository

import "github.com/gab-rod23/minitweeter/users/entities/model"

type UserRepository interface {
	InsertUser(newUser *model.UserModelCollection)
	FindUserByField(value string, field string) (*model.UserModelCollection, error)
}
