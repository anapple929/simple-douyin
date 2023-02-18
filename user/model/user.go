package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
	"time"
)

type User struct {
	UserId          int64  `gorm:"primary_key"`
	Name            string `gorm:"default:(-)"`
	FollowingCount  int64  `gorm:"default:(-)"`
	FollowerCount   int64  `gorm:"default:(-)"`
	Password        string `gorm:"default:(-)"`
	Avatar          string `gorm:"default:(-)"`
	BackgroundImage string `gorm:"default:(-)"`
	Signature       string `gorm:"default:(-)"`
	TotalFavorited  int64  `gorm:"default:(-)"`
	WorkCount       int64  `gorm:"default:(-)"`
	FavoriteCount   int64  `gorm:"default:(-)"`
	CreateAt        time.Time
	DeleteAt        time.Time
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once //单例模式，只生成一个userDao实例，提高性能

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

/**
根据用户名和密码，创建一个新的User，返回UserId
*/
func (*UserDao) CreateUser(user *User) (int64, error) {
	/*user := User{Name: username, Password: password, FollowingCount: 0, FollowerCount: 0, CreateAt: time.Now()}*/

	result := DB.Create(&user)

	if result.Error != nil {
		return -1, result.Error
	}

	return user.UserId, nil
}

/**
根据用户名，查找用户实体
*/
func (*UserDao) FindUserByName(username string) (*User, error) {
	user := User{Name: username}

	result := DB.Where("name = ?", username).First(&user)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

/**
根据用户id，查找用户实体
*/
func (d *UserDao) FindUserById(id int64) (*User, error) {
	user := User{UserId: id}

	result := DB.Where("user_id = ?", id).First(&user)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

/**
根据userId，给user表的follower字段增加cnt
*/
func (*UserDao) AddFollowerCount(userId int64, cnt int32) {
	err := DB.Model(User{}).Where("user_id=?", userId).Update("follower_count", gorm.Expr("follower_count+?", cnt)).Error
	if err != nil {
		//log.Error(err)
	}
}

/**
根据userId，给user表的follower字段减少cnt
*/
func (*UserDao) ReduceFollowerCount(userId int64, cnt int32) {
	err := DB.Model(User{}).Where("user_id=?", userId).Update("follower_count", gorm.Expr("follower_count-?", cnt)).Error
	if err != nil {
	}
}

/**
根据userId，给user表的following字段增加count
*/
func (*UserDao) AddFollowingCount(userId int64, count int32) {
	err := DB.Model(User{}).Where("user_id=?", userId).Update("following_count", gorm.Expr("following_count+?", count)).Error
	if err != nil {
		//log.Error(err)
	}
}

/**
根据userId，给user表的following字段减少cnt
*/
func (*UserDao) ReduceFollowingCount(userId int64, count int32) {
	err := DB.Model(User{}).Where("user_id=?", userId).Update("following_count", gorm.Expr("following_count-?", count)).Error
	if err != nil {
	}
}

/**
根据用户id集，获取user实体集
*/
func (*UserDao) GetUsersByIds(userIds []int64) ([]*User, error) {
	var users []*User

	err := DB.Where("user_id IN (?)", userIds).Find(&users).Error
	if err != nil {
		fmt.Println("model层查询User列表失败")
		return nil, err
	}

	return users, nil
}

/**
根据userId，给user表的work_count字段增加count
*/
func (*UserDao) AddWorkCount(userId int64, count int32) {
	err := DB.Model(User{}).Where("user_id=?", userId).Update("work_count", gorm.Expr("work_count+?", count)).Error
	if err != nil {
		//log.Error(err)
	}
}

/**
根据userId，给user表的work_count字段减少cnt
*/
func (*UserDao) ReduceWorkCount(userId int64, count int32) {
	err := DB.Model(User{}).Where("user_id=?", userId).Update("work_count", gorm.Expr("work_count-?", count)).Error
	if err != nil {
	}
}

/**
根据userId，给user表的favorite_count字段增加count
*/
func (*UserDao) AddFavoriteCount(userId int64, count int32) {
	err := DB.Model(User{}).Where("user_id=?", userId).Update("favorite_count", gorm.Expr("favorite_count+?", count)).Error
	if err != nil {
		//log.Error(err)
	}
}

/**
根据userId，给user表的favorite_count字段减少cnt
*/
func (*UserDao) ReduceFavoriteCount(userId int64, count int32) {
	err := DB.Model(User{}).Where("user_id=?", userId).Update("favorite_count", gorm.Expr("favorite_count-?", count)).Error
	if err != nil {
	}
}

/**
根据userId，给user表的total_favorited字段增加count
*/
func (*UserDao) AddTotalFavorited(userId int64, count int32) {
	err := DB.Model(User{}).Where("user_id=?", userId).Update("total_favorited", gorm.Expr("total_favorited+?", count)).Error
	if err != nil {
		//log.Error(err)
	}
}

/**
根据userId，给user表的total_favorited字段减少cnt
*/
func (*UserDao) ReduceTotalFavorited(userId int64, count int32) {
	err := DB.Model(User{}).Where("user_id=?", userId).Update("total_favorited", gorm.Expr("total_favorited-?", count)).Error
	if err != nil {
	}
}
