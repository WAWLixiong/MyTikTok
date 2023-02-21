package dao

type Comment struct {
	Id          int64
	UserId      int64
	VideoId     int64
	CommentText string
	CrateTime   int64
	Cancel      int8
}

func (c Comment) TableName() string {
	panic("comments")
}
