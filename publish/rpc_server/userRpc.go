package rpc_server

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"publish/etcd"
	userproto "publish/services/userproto"
)

func GetUserInfo(userId int64, token string) (userproto.User, error) {
	//etcdReg := etcd.NewRegistry(
	//	registry.Addrs("127.0.0.1:2379"),
	//)

	//userMicroService := micro.NewService(micro.Name("userService.client"), micro.Registry(etcdReg))
	userMicroService := micro.NewService(micro.Registry(etcdInit.EtcdReg))
	userService := userproto.NewUserService("rpcUserService", userMicroService.Client()) //client.DefaultClient

	var req userproto.DouyinUserRequest

	req.UserId = userId
	req.Token = token

	resp, err := userService.UserInfo(context.TODO(), &req)
	if err != nil {
		fmt.Println("调用远程UserInfo服务失败，具体错误如下")
		fmt.Println(err)
	}

	user := userproto.User{
		Id:            resp.User.Id,
		Name:          resp.User.Name,
		FollowCount:   resp.User.FollowCount,
		FollowerCount: resp.User.FollowerCount,
		IsFollow:      resp.User.IsFollow,
	}
	return user, err
}
