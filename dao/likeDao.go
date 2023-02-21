package dao

type Like struct {
	Id        int64
	UserId    int64
	VideoId   int64
	CrateTime int64
	Cancel    int8
}

func (l Like) TableName() string {
	panic("likes")
}
