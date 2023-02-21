package controller

import (
	"MyTikTok/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
}

func (u UserController) Login(ctx *gin.Context) {

}

func (u UserController) Logout(ctx *gin.Context) {

}

// UserInfo GET douyin/user/ 用户信息
func (u UserController) UserInfo(ctx *gin.Context) {
	user_id := ctx.Query("user_id")
	id, _ := strconv.ParseInt(user_id, 10, 64)

	user := service.UserServiceImpl{}
}
