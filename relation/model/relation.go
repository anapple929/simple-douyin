package model

import (
	"fmt"
	"sync"
	"time"
)

type Relation struct {
	FollowerId  int64
	FollowingId int64
	CreateAt    time.Time
}

func (Relation) TableName() string {
	return "relation"
}

type RelationDao struct {
}

func (d RelationDao) CreateRelation(relation *Relation) error {
	//和数据库进行操作
	result := DB.Create(&relation)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d RelationDao) DeleteRelation(relation *Relation) error {
	err := DB.Where("follower_id=? and following_id=?", relation.FollowerId, relation.FollowingId).Delete(relation).Error
	if err != nil {
		fmt.Println("取关失败")
		return err
	}
	return nil
}

func (d RelationDao) QueryFollowingIds(id int64) []int64 {
	var Ids []int64
	var results []*Relation
	DB.Select("follower_id").Where("following_id=?", id).Find(&results)

	for i := 0; i < len(results); i++ {
		Ids = append(Ids, results[i].FollowerId)
	}
	fmt.Println("查询到的关注的人的集合是：")
	fmt.Println(Ids)
	return Ids
}

func (d RelationDao) QueryFollowerIds(id int64) []int64 {
	var Ids []int64
	var results []*Relation
	DB.Select("following_id").Where("follower_id=?", id).Find(&results)
	for i := 0; i < len(results); i++ {
		Ids = append(Ids, results[i].FollowingId)
	}
	fmt.Println("查询到的粉丝集合是：")
	fmt.Println(Ids)
	return Ids
}

func (d RelationDao) QueryFriendIds(userId int64, ids []int64) []int64 {
	var Ids []int64
	var count int32
	for _, id := range ids {
		_ = DB.Model(&Relation{}).Where("following_id=? and follower_id=?", id, userId).Count(&count).Error
		if count > 0 {
			Ids = append(Ids, id)
		}
	}
	fmt.Println("输出一下朋友的id是什么")
	fmt.Println(Ids)
	return Ids
}

var relationDao *RelationDao
var relationOnce sync.Once //单例模式，只生成一个commentDao实例，提高性能

func NewRelationDaoInstance() *RelationDao {
	relationOnce.Do(
		func() {
			relationDao = &RelationDao{}
		})
	return relationDao
}
