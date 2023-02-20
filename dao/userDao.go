package dao

type User struct {
	Id       int64
	Name     string
	Password string
}

func (u User) TableName() string {
	panic("users")
}
