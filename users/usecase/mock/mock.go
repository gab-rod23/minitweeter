package mock

import (
	"github.com/gab-rod23/minitweeter/users/entities/dto"
	"github.com/stretchr/testify/mock"
)

type usecaseMock struct {
	mock *mock.Mock
}

type UserUsecaseMock interface {
	PatchCreateNewUser(newUserData *dto.CreateUserRequestDTO, expectedErr error)
	CreateNewUser(newUserData *dto.CreateUserRequestDTO) error
	PatchFollowUser(username string, followUserData *dto.FollowUserRequestDTO, expectedErr error)
	FollowUser(username string, followUserData *dto.FollowUserRequestDTO) error
	PatchRetrieveUserByUsername(username string, response *dto.UseDataResponseDTO, expectedErr error)
	RetrieveUserByUsername(username string) (*dto.UseDataResponseDTO, error)
}

func NewUserUsecaseMock() UserUsecaseMock {
	return &usecaseMock{
		mock: &mock.Mock{},
	}
}

func (m usecaseMock) PatchCreateNewUser(newUserData *dto.CreateUserRequestDTO, expectedErr error) {
	m.mock.On("CreateNewUser", newUserData).Return(expectedErr)
}

func (m usecaseMock) CreateNewUser(newUserData *dto.CreateUserRequestDTO) error {
	args := m.mock.Called(newUserData)
	return args.Error(0)
}

func (m usecaseMock) PatchFollowUser(username string, followUserData *dto.FollowUserRequestDTO, expectedErr error) {
	m.mock.On("FollowUser", username, followUserData).Return(expectedErr)
}

func (m usecaseMock) FollowUser(username string, followUserData *dto.FollowUserRequestDTO) error {
	args := m.mock.Called(username, followUserData)
	return args.Error(0)
}

func (m usecaseMock) PatchRetrieveUserByUsername(username string, response *dto.UseDataResponseDTO, expectedErr error) {
	m.mock.On("RetrieveUserByUsername", username, response).Return(expectedErr)
}

func (m usecaseMock) RetrieveUserByUsername(username string) (*dto.UseDataResponseDTO, error) {
	args := m.mock.Called(username)
	response := args.Get(0)
	if response == nil {
		return nil, args.Error(1)
	}
	return response.(*dto.UseDataResponseDTO), nil
}
