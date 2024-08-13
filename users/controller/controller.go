package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	HandlerCreateNewUser(ctx *gin.Context)
	HandlerRetrieveUserDataByUsername(ctx *gin.Context)
	HandlerFollowUser(ctx *gin.Context)
}
