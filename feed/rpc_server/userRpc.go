package rpc_server

import (
	"context"
	"feed/rpc_server/etcd"
	userproto "feed/services/userproto"
	"fmt"
	"github.com/micro/go-micro/v2"
)

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
