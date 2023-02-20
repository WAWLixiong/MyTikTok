package dao

type Video struct {
	Id       int64
	Title    string
	VideoUrl string
	CoverUrl string
	UserId   int64
}

func (v Video) TableName() string {
	panic("videos")
}
