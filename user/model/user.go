package model

import (
	"sync"
	"time"
)

type User struct {
	UserId         int64  `gorm:"primaryKey"`
	Name           string `gorm:"default:(-)"`
	FollowingCount int64  `gorm:"default:(-)"`
	FollowerCount  int64  `gorm:"default:(-)"`
	Password       string `gorm:"default:(-)"`
	CreateAt       time.Time
	DeleteAt       time.Time
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
