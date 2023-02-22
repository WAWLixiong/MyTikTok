package service

import (
	"MyTikTok/dao"
	"log"
)

type UserServiceImpl struct {
}

func (u UserServiceImpl) GetTableUserList() []dao.User {
	users, err := dao.GetUserList()
	if err != nil {
		log.Println("Err:", err.Error())
		return users
	}
	return users
}

func (u UserServiceImpl) GetTableUserByUsername(name string) dao.User {
	user, err := dao.GetUserByName(name)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User not found")
		return user
	}
	log.Println("Query user success")
	return user
}

func (u UserServiceImpl) GetTableUserById(id int64) dao.User {
	user, err := dao.GetUserById(id)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User not found")
		return user
	}
	log.Println("Query user success")
	return user
}

func (u UserServiceImpl) InsertTableUser(user *dao.User) bool {
	flag := dao.InsertUser(user)
	if flag == false {
		log.Println("插入失败")
		return false
	}
	return true
}

func (u UserServiceImpl) GetUserById(id int64) (User, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) GetUserByIdWithCurId(id int64, curId int64) (User, error) {
	// TODO implement me
	panic("implement me")
}
