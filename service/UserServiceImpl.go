package service

import (
	"MyTikTok/config"
	"MyTikTok/dao"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"strconv"
	"time"
)

type UserServiceImpl struct {
}

// GetUserList 查找所有用户
func (u UserServiceImpl) GetUserList() []dao.User {
	users, err := dao.GetUserList()
	if err != nil {
		log.Println("Err:", err.Error())
		return users
	}
	return users
}

// GetUserByUsername 通过用户名查找用户
func (u UserServiceImpl) GetUserByUsername(name string) dao.User {
	user, err := dao.GetUserByName(name)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User not found")
		return user
	}
	log.Println("Query user success")
	return user
}

// GetUserById 通过用户id查找用户
func (u UserServiceImpl) GetUserById(id int64) dao.User {
	user, err := dao.GetUserById(id)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User not found")
		return user
	}
	log.Println("Query user success")
	return user
}

// InsertTableUser 创建用户
func (u UserServiceImpl) InsertTableUser(user *dao.User) bool {
	flag := dao.InsertUser(user)
	if flag == false {
		log.Println("插入失败")
		return false
	}
	return true
}

// GetUserByIdNotLogin 通过id获取未登录状态的用户信息
func (u UserServiceImpl) GetUserByIdNotLogin(id int64) (User, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) GetUserByIdWithCurId(id int64, curId int64) (User, error) {
	// TODO implement me
	panic("implement me")
}

// EnCoder 密码加密
func EnCoder(password string) string {
	h := hmac.New(sha256.New, []byte(password))
	sha := hex.EncodeToString(h.Sum(nil))
	fmt.Println("Result: " + sha)
	return sha
}

// GenerateToken 根据username生成一个token
func GenerateToken(username string) string {
	user := UserService.GetUserByUsername(new(UserServiceImpl), username)
	fmt.Printf("generatetoken: %v\n", user)
	token := NewToken(user)
	return token
}

func NewToken(user dao.User) string {
	expiresTime := time.Now().Unix() + int64(config.OneDayOfHours)
	fmt.Printf("expiresTime: %v\n", expiresTime)
	id64 := user.Id
	fmt.Printf("id: %v\n", strconv.FormatInt(id64, 10))
	claims := jwt.StandardClaims{
		Audience:  user.Name,
		ExpiresAt: expiresTime,
		Id:        strconv.FormatInt(id64, 10),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "tiktok",
		NotBefore: time.Now().Unix(),
		Subject:   "token",
	}
	var jwtSecret = []byte(config.Secret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
		token = "Bearer" + token
		println("generate token success!\n")
		return token
	} else {
		return "fail"
	}

}
