package router

import "github.com/gin-gonic/gin"

func RoutersInit(engine *gin.Engine) {
	UserRouterInit(engine)
}
