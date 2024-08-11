package impl

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gab-rod23/minitweeter/users/controller"
	"github.com/gab-rod23/minitweeter/users/entities"
	"github.com/gab-rod23/minitweeter/users/usecase"
	"github.com/gab-rod23/minitweeter/users/usecase/impl"
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
	createUserRequestDto := new(entities.CreateUserRequestDTO)
	var err error

	err = ctx.BindJSON(&createUserRequestDto)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	err = u.userUsecase.CreateNewUser(createUserRequestDto)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusCreated, nil)
}

func (u userController) HandlerFollowUser(ctx *gin.Context) {
	followUserRequestDto := new(entities.FollowUserRequestDTO)
	var err error

	err = ctx.BindJSON(&followUserRequestDto)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	err = u.userUsecase.FollowUser(followUserRequestDto)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, nil)
}

func (u userController) HandlerRetrieveUserData(ctx *gin.Context) {
	username := ctx.GetHeader("username")
	if len(username) == 0 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Se debe enviar un usuario valido"))
	}
	userData, err := u.userUsecase.RetrieveUserByUsername(username)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, userData)
}

func (u userController) HandlerRetrieveUserFollowData(ctx *gin.Context) {

}
