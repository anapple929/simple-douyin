package model

import "time"

type Favorite struct {
	UserId   int64
	VideoId  int64
	CreateAt time.Time
}

func (Favorite) TableName() string {
	return "favorite"
}
