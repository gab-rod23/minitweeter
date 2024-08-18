package mock

import (
	"github.com/gab-rod23/minitweeter/tweets/entities/dto"
	"github.com/gab-rod23/minitweeter/tweets/entities/model"
	"github.com/gab-rod23/minitweeter/tweets/repository"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock *mock.Mock
}

func NewTweetRepositoryMock(m *mock.Mock) repository.TweetRepository {
	return &repositoryMock{
		mock: m,
	}
}

func (m repositoryMock) PatchInsertTweet(newTweet *model.TweetModelCollection, expectedErr error) {
	m.mock.On("InsertTweet", newTweet).Return(expectedErr)
}

func (m repositoryMock) InsertTweet(newTweet *model.TweetModelCollection) error {
	args := m.mock.Called(newTweet)
	return args.Error(0)
}

func (m repositoryMock) PatchFindTweetsFromUsers(timelineData *dto.TimelineTweetData, followingUsers []string, expectedTweetModelCollection []model.TweetModelCollection, expectedErr error) {
	m.mock.On("FindTweetsFromUsers", timelineData, followingUsers).Return(expectedTweetModelCollection, expectedErr)
}
func (m repositoryMock) FindTweetsFromUsers(timelineData *dto.TimelineTweetData, followingUsers []string) ([]model.TweetModelCollection, error) {
	args := m.mock.Called(timelineData, followingUsers)
	response := args.Get(0)
	if response == nil {
		return nil, args.Error(1)
	}
	return response.([]model.TweetModelCollection), nil
}
