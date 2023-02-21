package service

import (
	"MyTikTok/dao"
)

type UserServiceImpl struct {
}

func (u UserServiceImpl) GetTableUserList() []dao.User {
	// TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) GetTableUserByUsername(name string) dao.User {
	// TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) GetTableUserById(id int64) dao.User {
	// TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) InsertTableUser(tableUser *dao.User) bool {
	// TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) GetUserById(id int64) (User, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) GetUserByIdWithCurId(id int64, curId int64) (User, error) {
	// TODO implement me
	panic("implement me")
}
