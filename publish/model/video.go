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
根据创建一个新的Video，返回Video实例
*/
func (*VideoDao) CreateVideo(video *Video) (*Video, error) {
	/*Video := Video{Name: Videoname, Password: password, FollowingCount: 0, FollowerCount: 0, CreateAt: time.Now()}*/

	result := DB.Create(&video)

	if result.Error != nil {
		return nil, result.Error
	}

	return video, nil
}

/**
根据videoid，查找用户实体
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
