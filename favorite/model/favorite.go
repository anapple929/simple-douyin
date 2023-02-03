package model

import (
	"time"
)

type Favorite struct {
	UserId    int64
	VideoId   int64
	CreatedAt time.Time
}

func (Favorite) TableName() string {
	return "favorite"
}
