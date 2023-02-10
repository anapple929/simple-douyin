package model

import (
	"sync"
	"time"
)

type Comment struct {
	CommentId int64  `gorm:"primaryKey"`
	UserId    int64  `gorm:"default:(-)"`
	VideoId   int64  `gorm:"default:(-)"`
	Content   string `gorm:"default:(-)"`
	CreateAt  time.Time
	DeleteAt  time.Time
}

func (Comment) TableName() string {
	return "comment"
}

type CommentDao struct {
}

/**
根据commentId删除信息，软删除，因为有deleteat字段,返回这个操作有错误吗
*/
func (*CommentDao) DeleteCommentById(commentId int64) error {
	return nil
}

/**
创建一条Comment,返回创建的comment和error信息
*/
func (*CommentDao) CreateComment(comment *Comment) (*Comment, error) {
	//和数据库进行操作
	return &Comment{
		CommentId: 1,
		UserId:    51,
		Content:   "第一条测试数据",
		VideoId:   40,
		CreateAt:  time.Now(),
	}, nil
}

/**
传入videoId，查出comments
*/
func (*CommentDao) QueryComment(videoId int64) []*Comment {
	return []*Comment{
		{
			CommentId: 1,
			UserId:    51,
			Content:   "第一条测试数据",
			VideoId:   videoId,
			CreateAt:  time.Now(),
			DeleteAt:  time.Now(),
		},
		{
			CommentId: 2,
			UserId:    52,
			Content:   "第2条测试数据",
			VideoId:   videoId,
			CreateAt:  time.Now(),
			DeleteAt:  time.Now(),
		},
	}
}

var commentDao *CommentDao
var commentOnce sync.Once //单例模式，只生成一个commentDao实例，提高性能

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(
		func() {
			commentDao = &CommentDao{}
		})
	return commentDao
}
