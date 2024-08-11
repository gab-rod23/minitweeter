package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	HandlerCreateNewUser(ctx *gin.Context)
	HandlerRetrieveUserData(ctx *gin.Context)
	HandlerRetrieveUserFollowData(ctx *gin.Context)
	HandlerFollowUser(ctx *gin.Context)
}
