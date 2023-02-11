package model

import (
	"sync"
	"time"
	to_user "user/services/from_relation"
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

func (RelationDao) GetRelationsByIds(relationStatus []*to_user.RelationStatus) ([]*to_user.RelationStatus, error) {
	var result []*to_user.RelationStatus
	//调用relation数据库，查出relationStatus每一项的isFollow，封装成result返回.

	//查数据库
	result = relationStatus
	return result, nil
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
根据follower_id和followering_id，查找是否存在关系，返回bool值
*/
func (d *UserDao) FindRelationById(followerId int64, followeringId int64) (bool, error) {
	var count int
	err := DB.Model(Relation{}).Where("follower_id=? and following_id=?", followerId, followeringId).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 1, nil
}
