package impl

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gab-rod23/minitweeter/users/controller"
	"github.com/gab-rod23/minitweeter/users/entities/dto"
	"github.com/gab-rod23/minitweeter/users/usecase"
	"github.com/gab-rod23/minitweeter/users/usecase/impl"
	"github.com/gab-rod23/minitweeter/util"
)

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController() controller.UserController {
	return &userController{
		userUsecase: impl.NewUserUsecase(),
	}
}

func (u userController) HandlerCreateNewUser(ctx *gin.Context) {
	createUserRequestDto := new(dto.CreateUserRequestDTO)
	var err error

	err = ctx.BindJSON(&createUserRequestDto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrInvalidRequest.Error())
	}
	err = u.userUsecase.CreateNewUser(createUserRequestDto)
	if err != nil {
		if err == util.ErrEmailAlreadyExists || err == util.ErrUsernameAlreadyExists {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}

func (u userController) HandlerFollowUser(ctx *gin.Context) {
	username := ctx.GetHeader("username")
	followUserRequestDto := new(dto.FollowUserRequestDTO)
	var err error

	err = ctx.BindJSON(&followUserRequestDto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrInvalidRequest.Error())
		return
	}

	err = u.userUsecase.FollowUser(username, followUserRequestDto)
	if err != nil {
		if errors.Is(err, util.ErrUserNotFound) || errors.Is(err, util.ErrUserToFollowNotFound) {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (u userController) HandlerRetrieveUserDataByUsername(ctx *gin.Context) {
	username := ctx.GetHeader("username")
	if len(username) == 0 {
		ctx.JSON(http.StatusBadRequest, util.ErrInvalidUser.Error())
		return
	}
	userData, err := u.userUsecase.RetrieveUserByUsername(username)
	if err != nil {
		if errors.Is(err, util.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, userData)
}
