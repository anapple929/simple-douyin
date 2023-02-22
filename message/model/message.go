package model

import (
	"fmt"
	"gorm.io/gorm"
	"message/service/to_relation"
	"sync"
	"time"
)

type Message struct {
	Id         int64 `gorm:"primary_key"`
	FromUserId int64
	ToUserId   int64
	Content    string
	CreateAt   time.Time
	DeletedAt  gorm.DeletedAt
}

func (Message) TableName() string {
	return "message"
}

type MessageDao struct {
}

var messageDao *MessageDao
var messageOnce sync.Once //单例模式，只生成一个commentDao实例，提高性能

func NewMessageDaoInstance() *MessageDao {
	messageOnce.Do(
		func() {
			messageDao = &MessageDao{}
		})
	return messageDao
}

/**
创建一条消息
*/
func (*MessageDao) CreateMessage(message *Message) error {
	//和数据库进行操作
	result := DB.Create(&message)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

/**
查询消息记录
*/
//根据时间和需要查询的条数，获取video列表
func (*MessageDao) QueryMessageList(date *string, fromUserId int64, ToUserId int64) []*Message {
	fmt.Println(*date)
	var MessageList []*Message
	DB.Where("( (from_user_id = ? and to_user_id = ?) or (from_user_id = ? and to_user_id = ?) ) and create_at > ?", fromUserId, ToUserId, ToUserId, fromUserId, date).Order("create_at asc").Find(&MessageList)

	fmt.Println(MessageList)
	return MessageList
}

/**

 */
func (d *MessageDao) QueryMessage(toUserId int64, FromUserId int64) *to_relation.QueryBody {
	message := Message{}
	var msgType int64
	err := DB.Model(&Message{}).Where("(to_user_id=? and from_user_id=?) or (to_user_id=? and from_user_id=?)", toUserId, FromUserId, FromUserId, toUserId).Order("create_at desc").First(&message).Error
	if err != nil { //没查到，first会报个错
		return &to_relation.QueryBody{
			FromUserId: FromUserId,
			ToUserId:   toUserId,
			Message:    &to_relation.Message{},
			MsgType:    0,
		}
	}
	if FromUserId == message.FromUserId {
		msgType = 1
	} else {
		msgType = 0
	}

	return &to_relation.QueryBody{
		FromUserId: FromUserId,
		ToUserId:   toUserId,
		Message: &to_relation.Message{
			Id:      message.Id,
			Content: message.Content,
		},
		MsgType: msgType,
	}
}
