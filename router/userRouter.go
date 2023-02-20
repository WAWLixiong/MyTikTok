package router

import (
	"MyTikTok/controller"
	"github.com/gin-gonic/gin"
)

func UserRouterInit(router *gin.Engine) {
	userRouter := router.Group("/user")
	{
		userRouter.GET("/user", controller.UserController{}.User)
		userRouter.GET("/login", controller.UserController{}.User)
		userRouter.GET("/logout", controller.UserController{}.User)
	}
}
