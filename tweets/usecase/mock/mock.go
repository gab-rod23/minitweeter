package mock

import (
	"github.com/gab-rod23/minitweeter/tweets/entities/dto"
	"github.com/gab-rod23/minitweeter/tweets/usecase"
	"github.com/stretchr/testify/mock"
)

type usecaseMock struct {
	mock *mock.Mock
}

func NewTweetUsecaseMock(m *mock.Mock) usecase.TweetUsecase {
	return &usecaseMock{
		mock: m,
	}
}

func (m usecaseMock) PatchCreateNewTweet(tweet *dto.CreateTweetRequestDto, username string, expectedErr error) {
	m.mock.On("CreateNewTweet", tweet, username).Return(expectedErr)
}

func (m usecaseMock) CreateNewTweet(tweet *dto.CreateTweetRequestDto, username string) error {
	args := m.mock.Called(tweet, username)
	return args.Error(0)
}

func (m usecaseMock) PatchRetrieveTimelineTweet(timelineData *dto.TimelineTweetData, expectedResponse *dto.TimelineTweetResponseDto, expectedErr error) {
	m.mock.On("CreateNewTweet", timelineData).Return(expectedResponse, expectedErr)
}

func (m usecaseMock) RetrieveTimelineTweet(timelineData *dto.TimelineTweetData) (*dto.TimelineTweetResponseDto, error) {
	args := m.mock.Called(timelineData)
	response := args.Get(0)
	if response == nil {
		return nil, args.Error(1)
	}
	return response.(*dto.TimelineTweetResponseDto), nil
}
