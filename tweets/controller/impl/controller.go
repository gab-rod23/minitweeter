package impl

import (
	"net/http"
	"strconv"

	"github.com/gab-rod23/minitweeter/tweets/controller"
	"github.com/gab-rod23/minitweeter/tweets/entities/dto"
	"github.com/gab-rod23/minitweeter/tweets/usecase"
	"github.com/gab-rod23/minitweeter/tweets/usecase/impl"
	"github.com/gin-gonic/gin"
)

type tweetController struct {
	tweetUsecase usecase.TweetUsecase
}

func NewTweetController() controller.TweetController {
	return &tweetController{
		tweetUsecase: impl.NewTweetUsecase(),
	}
}

func (t tweetController) HandlerCreateNewTweet(ctx *gin.Context) {
	username := ctx.GetHeader("username")
	createNewTweet := new(dto.CreateTweetRequestDto)
	err := ctx.BindJSON(createNewTweet)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	err = t.tweetUsecase.CreateNewTweet(createNewTweet, username)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusCreated, nil)
}

func (t tweetController) HandlerRetrieveTimelineTweet(ctx *gin.Context) {
	timelineTweetData, err := validateAndGenerateTimelineTweetData(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	timeline, err := t.tweetUsecase.RetrieveTimelineTweet(timelineTweetData)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, timeline)
}

func validateAndGenerateTimelineTweetData(ctx *gin.Context) (*dto.TimelineTweetData, error) {
	username := ctx.GetHeader("username")
	pageSize := ctx.GetHeader("page_size")
	pageNumber := ctx.GetHeader("page_number")
	//	lastPageTweetDate := ctx.GetHeader("last_page_tweet_date")
	//	parsedTime, _ := time.Parse(time.RFC3339, lastPageTweetDate)
	parsedPageSize, _ := strconv.Atoi(pageSize)
	parsedPageNumber, _ := strconv.Atoi(pageNumber)
	return &dto.TimelineTweetData{
		Username:   username,
		PageSize:   parsedPageSize,
		PageNumber: parsedPageNumber,
	}, nil
}
