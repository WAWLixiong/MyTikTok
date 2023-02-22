package dao

import "log"

type User struct {
	Id       int64
	Name     string
	Password string
}

func (u User) TableName() string {
	return "users"
}

// GetUserList 获取全部User对象
func GetUserList() ([]User, error) {
	users := []User{}
	if err := Db.Find(&users).Error; err != nil {
		log.Println(err.Error())
		return users, err
	}
	return users, nil
}

// GetUserByName 通过名称获取User对象
func GetUserByName(name string) (User, error) {
	user := User{}
	if err := Db.Where("name = ?", name).First(&user).Error; err != nil {
		log.Println(err.Error())
		return user, err
	}
	return user, nil
}

// GetUserById 通过id获取User对象
func GetUserById(id int64) (User, error) {
	user := User{}
	if err := Db.Where("id = ?", id).First(&user).Error; err != nil {
		log.Println(err.Error())
		return user, err
	}
	return user, nil
}

func InsertUser(user *User) bool {
	if err := Db.Create(user).Error; err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
