package rpc_server

import (
	"comment/rpc_server/etcd"
	usersproto "comment/service/to_relation"
	userproto "comment/service/userproto"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
)

/**
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

/**
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
