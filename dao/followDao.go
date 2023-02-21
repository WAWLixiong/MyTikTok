package dao

type Follow struct {
	Id       int64
	UserId   int64
	FollowId int64
	Cancel   int8
}

func (f Follow) TableName() string {
	panic("follows")
}
