package rpc_server

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"relation/rpc_server/etcd"
	usersproto "relation/service/to_relation"
	userproto "relation/service/userproto"
)

/*
*
调用user查询用户信息
*/
func GetUserInfo(userId int64, token string) (*userproto.User, error) {
	userMicroService := micro.NewService(micro.Registry(etcdInit.EtcdReg))
	userService := userproto.NewUserService("rpcUserService", userMicroService.Client())

	var req userproto.DouyinUserRequest

	req.UserId = userId
	req.Token = token

	resp, err := userService.UserInfo(context.TODO(), &req)
	if err != nil {
		fmt.Println("调用远程UserInfo服务失败，具体错误如下")
		fmt.Println(err)
	}

	user := &userproto.User{
		Id:            resp.User.Id,
		Name:          resp.User.Name,
		FollowCount:   resp.User.FollowCount,
		FollowerCount: resp.User.FollowerCount,
		IsFollow:      resp.User.IsFollow,
	}
	return user, err
}

/*
*
输入userId列表，查询User实体列表
*/
func GetUsersInfo(userId []int64, token string) ([]*usersproto.User, error) {
	userMicroService := micro.NewService(micro.Registry(etcdInit.EtcdReg))
	usersService := usersproto.NewToRelationService("rpcUserService", userMicroService.Client())

	var req usersproto.GetUsersByIdsRequest

	req.UserId = userId
	req.Token = token

	resp, err := usersService.GetUsersByIds(context.TODO(), &req)
	if err != nil {
		fmt.Println("调用远程UserInfo服务失败，具体错误如下")
		fmt.Println(err)
	}
	fmt.Println("调用回来了")
	fmt.Println(resp.UserList)

	return resp.UserList, err
}

/*
*
更新user表的following_count字段
*/
func UpdateFollowingCount(followId int64, followerId int64, count int32, actionType int32) bool {
	MicroService := micro.NewService(micro.Registry(etcdInit.EtcdReg))
	Service := usersproto.NewToRelationService("rpcUserService", MicroService.Client())
	var req usersproto.UpdateFollowingCountRequest
	req.UserId = followId
	req.Count = count
	req.Type = actionType
	resp, err := Service.UpdateFollowingCount(context.TODO(), &req)
	if err != nil || resp.StatusCode != 0 {
		fmt.Println("followingCount维护失败:", err)
		return false
	}
	return true
}

/*
*
更新user表的follower_count字段
*/
func UpdateFollowerCount(followId int64, followerId int64, count int32, actionType int32) bool {
	MicroService := micro.NewService(micro.Registry(etcdInit.EtcdReg))
	Service := usersproto.NewToRelationService("rpcUserService", MicroService.Client())
	var req usersproto.UpdateFollowerCountRequest
	req.UserId = followerId
	req.Count = count
	req.Type = actionType
	resp, err := Service.UpdateFollowerCount(context.TODO(), &req)
	if err != nil || resp.StatusCode != 0 {
		fmt.Println("followerCount维护失败:", err)
		return false
	}
	return true
}
