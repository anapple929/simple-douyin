package core

import (
	"context"
	"fmt"
	"message/model"
	"message/rpc_server"
	proto "message/service"
	"strconv"
	"time"
)

type MessageService struct {
}

/**
消息列表
*/
func (*MessageService) MessageList(ctx context.Context, in *proto.DouyinMessageChatRequest, out *proto.DouyinMessageChatResponse) error {
	fmt.Println("调用了消息列表功能")
	//1.判断一下token失效了吗，调用rpc_server的 GetIdByToken 方法，从token中解析出userId
	//解析token
	if in.Token == "" {
		out.StatusCode = -1
		out.StatusMsg = "登录失效，请重新登录"
		return nil
	}
	userId, err := rpc_server.GetIdByToken(in.Token) //当前用户id
	if err != nil {
		out.StatusCode = -1
		out.StatusMsg = "登录失效，请重新登录"
		return nil
	}

	toUserId := in.ToUserId   //对方用户id
	lastTime := in.PreMsgTime //上次最新消息时间，毫秒。
	format := "2006-01-02 15:04:05"
	t := time.Unix(lastTime/1000, 0)
	searchTime := t.Format(format) //转化完的时间，可以去数据库查的数据

	var messageList []*model.Message //查出来的数据库列表
	var resultList []*proto.Message
	messageList = model.NewMessageDaoInstance().QueryMessageList(&searchTime, userId, toUserId)
	for _, message := range messageList {
		resultList = append(resultList, BuildProtoMessage(message))
	}

	out.MessageList = resultList
	out.StatusMsg = "查询消息成功"
	out.StatusCode = 0
	return nil
}

/**
发送消息
*/
func (*MessageService) MessageAction(ctx context.Context, in *proto.DouyinMessageActionRequest, out *proto.DouyinMessageActionResponse) error {
	fmt.Println("调用了发消息功能")
	//1.判断一下token失效了吗，调用rpc_server的 GetIdByToken 方法，从token中解析出userId
	//解析token
	if in.Token == "" {
		out.StatusCode = -1
		out.StatusMsg = "登录失效，请重新登录"
		return nil
	}
	userId, err := rpc_server.GetIdByToken(in.Token) //当前用户id
	if err != nil {
		out.StatusCode = -1
		out.StatusMsg = "登录失效，请重新登录"
		return nil
	}
	message := model.Message{
		FromUserId: userId,
		ToUserId:   in.ToUserId,
		Content:    in.Content,
		CreateAt:   time.Now(),
	}
	_ = model.NewMessageDaoInstance().CreateMessage(&message)

	out.StatusMsg = "发消息成功"
	out.StatusCode = 0
	return nil
}

func BuildProtoMessage(message *model.Message) *proto.Message {
	return &proto.Message{
		Id:         message.Id,
		ToUserId:   message.ToUserId,
		FromUserId: message.FromUserId,
		Content:    message.Content,
		CreateTime: strconv.FormatInt(message.CreateAt.UnixNano()/1e6, 10),
	}
}
