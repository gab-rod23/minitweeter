package impl

import (
	"net/http"

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

}
