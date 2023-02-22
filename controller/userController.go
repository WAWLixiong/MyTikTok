package controller

import (
	"MyTikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User service.User `json:"user"`
}

type UserController struct {
}

// Login POST douyin/user/login 用户登录
func (u UserController) Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	encoderPassword := service.EnCoder(password)
	println(encoderPassword)

	usi := service.UserServiceImpl{}

	user := usi.GetUserByUsername(username)
	if encoderPassword == user.Password {
		token := service.GenerateToken(username)
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Username or Password Error"},
		})
	}

}

func (u UserController) Logout(ctx *gin.Context) {

}

// UserInfo GET douyin/user/ 用户信息
func (u UserController) UserInfo(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	id, _ := strconv.ParseInt(userId, 10, 64)
	userService := service.UserServiceImpl{}
	user, err := userService.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusOK, UserResponse{Response: Response{StatusCode: 1, StatusMsg: "User Doesn't Exists"}})
	} else {
		ctx.JSON(http.StatusOK, UserResponse{Response: Response{StatusCode: 0}, User: user})
	}
}
