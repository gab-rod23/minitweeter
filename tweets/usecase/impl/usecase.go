package impl

import (
	"time"

	"github.com/gab-rod23/minitweeter/tweets/entities/dto"
	"github.com/gab-rod23/minitweeter/tweets/entities/model"
	"github.com/gab-rod23/minitweeter/tweets/repository"
	"github.com/gab-rod23/minitweeter/tweets/repository/impl"
	"github.com/gab-rod23/minitweeter/tweets/usecase"
	userRepo "github.com/gab-rod23/minitweeter/users/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type tweetUsecase struct {
	tweetRepository repository.TweetRepository
	userRepository  userRepo.UserRepository
}

func NewTweetUsecase() usecase.TweetUsecase {
	return &tweetUsecase{
		tweetRepository: impl.NewTweetRepository(),
	}
}

func (t tweetUsecase) CreateNewTweet(newTweetData *dto.CreateTweetRequestDto, username string) error {
	tweetToInsert := generateTweet(newTweetData, username)
	t.tweetRepository.InsertTweet(tweetToInsert)
	return nil
}

func (t tweetUsecase) RetrieveTimelineTweet(*dto.TimelineTweetData) (*dto.TimelineTweetResponseDto, error) {

	return nil, nil
}

func generateTweet(newUserData *dto.CreateTweetRequestDto, username string) *model.TweetModelCollection {
	return &model.TweetModelCollection{
		ID:          primitive.NewObjectID(),
		Username:    username,
		Text:        newUserData.Text,
		CreatedDate: primitive.NewDateTimeFromTime(time.Now()),
	}
}
