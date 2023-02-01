package model

type Favorite struct {
	UserId   string
	VideoId  string
	CreateAt string
}

func (Favorite) TableName() string {
	return "favorite"
}
