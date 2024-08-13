package impl

import (
	"time"

	"github.com/devfeel/mapper"
	"github.com/gab-rod23/minitweeter/tweets/entities/dto"
	"github.com/gab-rod23/minitweeter/tweets/entities/model"
	"github.com/gab-rod23/minitweeter/tweets/repository"
	"github.com/gab-rod23/minitweeter/tweets/repository/impl"
	"github.com/gab-rod23/minitweeter/tweets/usecase"
	userRepo "github.com/gab-rod23/minitweeter/users/repository"
	userRepoImpl "github.com/gab-rod23/minitweeter/users/repository/impl"
	"github.com/gab-rod23/minitweeter/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type tweetUsecase struct {
	tweetRepository repository.TweetRepository
	userRepository  userRepo.UserRepository
}

func NewTweetUsecase() usecase.TweetUsecase {
	return &tweetUsecase{
		tweetRepository: impl.NewTweetRepository(),
		userRepository:  userRepoImpl.NewUserRepository(),
	}
}

func (t tweetUsecase) CreateNewTweet(newTweetData *dto.CreateTweetRequestDto, username string) error {
	tweetToInsert := generateTweet(newTweetData, username)
	err := t.tweetRepository.InsertTweet(tweetToInsert)
	return err
}

func (t tweetUsecase) RetrieveTimelineTweet(timelineData *dto.TimelineTweetData) (*dto.TimelineTweetResponseDto, error) {
	userData, err := t.userRepository.FindUserByField(timelineData.Username, "username")
	if err != nil {
		return nil, util.ErrUserNotFound
	}
	timelineResult, err := t.tweetRepository.FindTweetsFromUsers(timelineData, userData.Following)
	if err != nil {
		return nil, err
	}
	timelineResponse := generateTimelineResponse(timelineResult)
	return timelineResponse, nil
}

func generateTweet(newUserData *dto.CreateTweetRequestDto, username string) *model.TweetModelCollection {
	return &model.TweetModelCollection{
		ID:          primitive.NewObjectID(),
		Username:    username,
		Text:        newUserData.Text,
		CreatedDate: time.Now(),
	}
}

func generateTimelineResponse(timelineModel []model.TweetModelCollection) *dto.TimelineTweetResponseDto {
	m := mapper.NewMapper()
	tweetArray := make([]dto.TimelineTweetResponseItem, 0)
	for _, item := range timelineModel {
		itemResponse := &dto.TimelineTweetResponseItem{}
		m.Mapper(&item, itemResponse)

		tweetArray = append(tweetArray, *itemResponse)
	}

	return &dto.TimelineTweetResponseDto{
		Timeline: tweetArray,
	}
}
