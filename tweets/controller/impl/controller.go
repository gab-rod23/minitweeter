package impl

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gab-rod23/minitweeter/tweets/controller"
	"github.com/gab-rod23/minitweeter/tweets/entities/dto"
	"github.com/gab-rod23/minitweeter/tweets/usecase"
	"github.com/gab-rod23/minitweeter/tweets/usecase/impl"
	"github.com/gab-rod23/minitweeter/util"
	"github.com/gin-gonic/gin"
)

const TWEET_LENGTH = 280

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
	if len(createNewTweet.Text) > TWEET_LENGTH {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("el tweet no puede superar los %d caracteres", TWEET_LENGTH))
		return
	}
	err := ctx.BindJSON(createNewTweet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrInvalidRequest.Error())
		return
	}
	err = t.tweetUsecase.CreateNewTweet(createNewTweet, username)
	if err != nil {
		if errors.Is(err, util.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, fmt.Sprintf("usuario %s: %s", username, err.Error()))
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}

func (t tweetController) HandlerRetrieveTimelineTweet(ctx *gin.Context) {
	timelineTweetData, err := validateAndGenerateTimelineTweetData(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	timeline, err := t.tweetUsecase.RetrieveTimelineTweet(timelineTweetData)
	if err != nil {
		if errors.Is(err, util.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, fmt.Sprintf("usuario %s: %s", timelineTweetData.Username, err.Error()))
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, timeline)
}

func validateAndGenerateTimelineTweetData(ctx *gin.Context) (*dto.TimelineTweetData, error) {
	var pageSize, pageNumber int
	var err error

	username := ctx.GetHeader("username")
	if len(username) == 0 {
		return nil, util.ErrInvalidUser
	}

	pageSize, err = strconv.Atoi(ctx.GetHeader("page_size"))
	fmt.Println(pageSize)
	if err != nil || pageSize <= 0 {
		return nil, util.ErrInvalidPageSize
	}

	pageNumberHeader := ctx.GetHeader("page_number")
	if len(pageNumberHeader) == 0 {
		pageNumber = 0
	} else {
		pageNumber, err = strconv.Atoi(pageNumberHeader)
		if err != nil || pageNumber < 0 {
			return nil, util.ErrInvalidPageNumber
		}
	}
	var lastPageTweetDate *time.Time
	lastPageTweetDateHeader := ctx.GetHeader("last_page_tweet_date")
	if len(lastPageTweetDateHeader) > 0 {
		parsedDate, err := time.Parse(time.RFC3339, lastPageTweetDateHeader)
		if err != nil {
			return nil, util.ErrInvalidDateFormat
		}
		lastPageTweetDate = &parsedDate
	}

	return &dto.TimelineTweetData{
		Username:          username,
		PageSize:          pageSize,
		PageNumber:        pageNumber,
		LastPageTweetDate: lastPageTweetDate,
	}, nil
}
