package core

import (
	"context"
	"fmt"
	"relation/model"
	"relation/rpc_server"
	proto "relation/service"
	from_message "relation/service/from_message"
	usersproto "relation/service/to_relation"
	"time"
)

type RelationService struct {
}

/*
*
关注
*/
func (*RelationService) RelationAction(ctx context.Context, in *proto.DouyinRelationActionRequest, out *proto.DouyinRelationActionResponse) error {
	//1.判断一下token失效了吗，调用rpc_server的 GetIdByToken 方法，从token中解析出userId
	//解析token
	if in.Token == "" {
		out.StatusCode = -1
		out.StatusMsg = "登录失效，请重新登录"
		return nil
	}
	userId, err := rpc_server.GetIdByToken(in.Token)
	if err != nil {
		out.StatusCode = -1
		out.StatusMsg = "登录失效，请重新登录"
		return nil
	}
	//token没有失效，解析出userID后拿到user控制层实体，把userID和token作为参数，调用rpc_server的GetUserInfo，查到User信息
	user, _ := rpc_server.GetUserInfo(userId, in.Token)
	fmt.Println(user) //输出看一下，查出来了吗

	//创建comment实体
	relation := &model.Relation{
		FollowerId:  in.ToUserId,
		FollowingId: userId,
		CreateAt:    time.Now(),
	}

	//判断一下actionType的类型 1 关注 2 取消关注， 做一个if判断
	if in.ActionType == 1 {
		//数据库插入数据
		_ = model.NewRelationDaoInstance().CreateRelation(relation)
		//调用rpc_server的user更新user微服务提供的方法，更新两个字段
		rpc_server.UpdateFollowingCount(userId, in.ToUserId, 1, in.ActionType)
		rpc_server.UpdateFollowerCount(userId, in.ToUserId, 1, in.ActionType)

		out.StatusMsg = "关注成功"
		out.StatusCode = 0
		return nil

	} else if in.ActionType == 2 {
		//数据库删除数据
		_ = model.NewRelationDaoInstance().DeleteRelation(relation)
		//调用rpc_server的user更新user微服务提供的方法，更新两个字段
		rpc_server.UpdateFollowingCount(userId, in.ToUserId, 1, in.ActionType)
		rpc_server.UpdateFollowerCount(userId, in.ToUserId, 1, in.ActionType)
		out.StatusMsg = "取关成功"
		out.StatusCode = 0
	} else {
		out.StatusCode = -1
		out.StatusMsg = "actionType有问题"
		return nil
	}

	return nil
}

/*
*
关注列表
*/
func (*RelationService) FollowList(ctx context.Context, in *proto.DouyinRelationFollowListRequest, out *proto.DouyinRelationFollowListResponse) error {
	//1.判断一下token失效了吗，调用rpc_server的 GetIdByToken 方法，从token中解析出userId
	//解析token
	if in.Token == "" {
		out.StatusCode = -1
		out.StatusMsg = "登录失效，请重新登录"
		return nil
	}
	_, err := rpc_server.GetIdByToken(in.Token)
	if err != nil {
		out.StatusCode = -1
		out.StatusMsg = "登录失效，请重新登录"
		return nil
	}
	//存评论列表
	var userResult []*proto.User
	//拿到userIds集合，调用usersinfo方法，查一批User实体
	var userIds []int64
	//调用数据库方法
	userIds = model.NewRelationDaoInstance().QueryFollowingIds(in.UserId)

	//调用usersinfo方法，查一批User实体
	users, _ := rpc_server.GetUsersInfo(userIds, in.Token)
	fmt.Println(users)
	//封装成视图层对象
	for _, user := range users {
		userResult = append(userResult, BuildProtoUser(user))
	}

	fmt.Println("输出一下朋友的信息")
	for _, user := range users {
		fmt.Println(user)
	}
	out.UserList = userResult
	out.StatusCode = 0
	out.StatusMsg = "查询关注列表成功"
	fmt.Println(out.UserList)
	return nil
}

/*
*
粉丝列表
*/
func (*RelationService) FollowerList(ctx context.Context, in *proto.DouyinRelationFollowerListRequest, out *proto.DouyinRelationFollowerListResponse) error {
	//1.判断一下token失效了吗，调用rpc_server的 GetIdByToken 方法，从token中解析出userId
	//解析token
	if in.Token == "" {
		out.StatusCode = -1
		out.StatusMsg = "登录失效，请重新登录"
		return nil
	}
	_, err := rpc_server.GetIdByToken(in.Token)
	if err != nil {
		out.StatusCode = -1
		out.StatusMsg = "登录失效，请重新登录"
		return nil
	}
	//存user列表
	var userResult []*proto.User
	//拿到userIds集合，调用usersinfo方法，查一批User实体
	var userIds []int64
	//调用数据库方法
	userIds = model.NewRelationDaoInstance().QueryFollowerIds(in.UserId)

	//调用usersinfo方法，查一批User实体
	users, _ := rpc_server.GetUsersInfo(userIds, in.Token)
	fmt.Println(users)
	//封装成视图层对象
	for _, user := range users {
		userResult = append(userResult, BuildProtoUser(user))
	}

	out.UserList = userResult
	out.StatusCode = 0
	out.StatusMsg = "查询粉丝列表成功"
	fmt.Println(out.UserList)
	return nil
}

/*
*
好友列表
*/
func (*RelationService) FriendList(ctx context.Context, in *proto.DouyinRelationFriendListRequest, out *proto.DouyinRelationFriendListResponse) error {
	//1.判断一下token失效了吗，调用rpc_server的 GetIdByToken 方法，从token中解析出userId
	//解析token
	if in.Token == "" {
		out.StatusCode = -1
		out.StatusMsg = "登录失效，请重新登录"
		return nil
	}
	userId, err := rpc_server.GetIdByToken(in.Token)
	if err != nil {
		out.StatusCode = -1
		out.StatusMsg = "登录失效，请重新登录"
		return nil
	}
	//userId是本人，查一下他的好友，先查他的粉丝id
	var userIds []int64
	//调用数据库方法,查询他的粉丝ids
	userIds = model.NewRelationDaoInstance().QueryFollowerIds(in.UserId)
	//查一下，粉丝id作为发起关注人，他作为被关注人，是否有关注信息，有就把这种粉丝id列表返回
	friendIds := model.NewRelationDaoInstance().QueryFriendIds(userId, userIds)
	//批量远程调用，根据朋友id列表，查询对应的Message列表
	//先构造批量调用的参数
	var queryBodys []*from_message.QueryBody
	for _, friendId := range friendIds {
		queryBodys = append(queryBodys, &from_message.QueryBody{
			FromUserId: userId,
			ToUserId:   friendId,
			Message:    nil,
			MsgType:    0,
		})
	}
	queryBodysResult, err := rpc_server.QueryMessagesByUsers(queryBodys)
	//通过朋友id列表，调用user，查到user信息，返回
	friendUserList, _ := rpc_server.GetUsersInfo(friendIds, in.Token)
	var resultList []*proto.FriendUser

	for _, friendUser := range friendUserList {
		//在批量返回的结果里，看一下，如果toUserId相等于朋友，那么就把这个queryBody取出来，封装最终结果
		for _, queryBodyResult := range queryBodysResult {
			if queryBodyResult.ToUserId == friendUser.Id {
				resultList = append(resultList, BuildProtoFriendUser(friendUser, queryBodyResult))
				break
			}
		}
	}
	out.StatusMsg = "查询好友成功"
	out.StatusCode = 0
	out.UserList = resultList
	return nil
}

func BuildProtoFriendUser(user *usersproto.User, queryBody *from_message.QueryBody) *proto.FriendUser {
	return &proto.FriendUser{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
		MsgType:       queryBody.MsgType,
		Message:       queryBody.Message.Content,
	}
}

/*
*
对user进行封装
*/
func BuildProtoUser(user *usersproto.User) *proto.User {
	return &proto.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}
