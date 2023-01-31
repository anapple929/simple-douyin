package model

import (
	"sync"
	"time"
)

type Relation struct {
	FollowerId  int64 `gorm:"default:(-)"`
	FollowingId int64 `gorm:"default:(-)"`
	CreateAt    time.Time
}

func (Relation) TableName() string {
	return "relation"
}

type RelationDao struct {
}

var relationDao *RelationDao
var relationOnce sync.Once //单例模式，只生成一个userDao实例，提高性能

func NewRelationDaoInstance() *RelationDao {
	relationOnce.Do(
		func() {
			relationDao = &RelationDao{}
		})
	return relationDao
}

/**
根据follower_id和followering_id，查找用户实体
*/
func (d *UserDao) FindRelationById(followerId int64, followeringId int64) (*Relation, error) {
	relation := Relation{FollowerId: followerId, FollowingId: followeringId}

	result := DB.Where("follower_id = ? and following_id = ?", followerId, followeringId).First(&relation)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &relation, err
}
