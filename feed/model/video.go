package model

import (
	"fmt"
	"sync"
	"time"
)

type Video struct {
	VideoId       int64  `gorm:"primaryKey"`
	UserId        int64  `gorm:"default:(-)"`
	PlayUrl       string `gorm:"default:(-)"`
	CoverUrl      string `gorm:"default:(-)"`
	FavoriteCount int64  `gorm:"default:(-)"`
	CommentCount  int64  `gorm:"default:(-)"`
	Title         string `gorm:"default:(-)"`
	CreateAt      time.Time
	UpdateAt      time.Time
	DeleteAt      time.Time
}

func (Video) TableName() string {
	return "video"
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once //单例模式，只生成一个VideoDao实例，提高性能

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

/**
根据videoid，查找video实体
*/
func (d *VideoDao) FindVideoById(id int64) (*Video, error) {
	video := Video{VideoId: id}

	result := DB.Where("Video_id = ?", id).First(&video)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &video, err
}

//根据UserId，查出Video列表
func (*VideoDao) QueryVideoByUserId(userId int64) ([]*Video, error) {
	var videos []*Video
	err := DB.Where("user_id = ?", userId).Find(&videos).Error
	if err != nil {
		fmt.Println("查询Video列表失败")
		return nil, err
	}
	return videos, nil
}

//根据时间和需要查询的条数，获取video列表
func (*VideoDao) QueryVideo(date *string, limit int) []*Video {
	fmt.Println(*date)
	var VideoList []*Video
	DB.Where("create_at < ?", *date).Order("create_at desc").Find(&VideoList)
	if len(VideoList) <= limit {
		fmt.Println(VideoList)
		return VideoList
	}
	fmt.Println(VideoList)
	return VideoList[0:limit]
}
