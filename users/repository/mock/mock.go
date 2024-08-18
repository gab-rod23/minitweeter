package mock

import (
	"github.com/gab-rod23/minitweeter/users/entities/model"
	"github.com/gab-rod23/minitweeter/users/repository"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock *mock.Mock
}

func NewUserRepositoryMock(m *mock.Mock) repository.UserRepository {
	return &repositoryMock{
		mock: m,
	}
}

func (m repositoryMock) PatchInsertUser(newUser *model.UserModelCollection, expectedErr error) {
	m.mock.On("FindUserByField", newUser).Return(expectedErr)
}

func (m repositoryMock) InsertUser(newUser *model.UserModelCollection) error {
	args := m.mock.Called(newUser)
	return args.Error(0)
}

func (m repositoryMock) PatchFindUserByField(value string, field string, expectedUserModelCollection *model.UserModelCollection, expectedErr error) {
	m.mock.On("FindUserByField", value, field, expectedUserModelCollection, expectedErr)
}

func (m repositoryMock) FindUserByField(value string, field string) (*model.UserModelCollection, error) {
	args := m.mock.Called(value, field)
	response := args.Get(0)
	if response == nil {
		return nil, args.Error(1)
	}
	return response.(*model.UserModelCollection), nil
}

func (m repositoryMock) PatchAddNewFollowerToUser(valueFilter string, fieldFilter string, followerUsername string, expectedErr error) {
	m.mock.On("AddNewFollowerToUser", valueFilter, fieldFilter, followerUsername).Return(expectedErr)
}

func (m repositoryMock) AddNewFollowerToUser(valueFilter string, fieldFilter string, followerUsername string) error {
	args := m.mock.Called(valueFilter, fieldFilter, followerUsername)
	return args.Error(0)
}

func (m repositoryMock) PatchAddNewFollowingToUser(valueFilter string, fieldFilter string, followingUsername string, expectedErr error) {
	m.mock.On("AddNewFollowingToUser", valueFilter, fieldFilter, followingUsername).Return(expectedErr)
}

func (m repositoryMock) AddNewFollowingToUser(valueFilter string, fieldFilter string, followingUsername string) error {
	args := m.mock.Called(valueFilter, fieldFilter, followingUsername)
	return args.Error(0)
}
