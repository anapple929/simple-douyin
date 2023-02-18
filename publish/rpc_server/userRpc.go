package rpc_server

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"publish/rpc_server/etcd"
	from_user_proto "publish/services/from_user"
	usersproto "publish/services/to_relation"
	userproto "publish/services/userproto"
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

	return resp.UserList, err
}

func UpdateWorkCount(uid int64, count int32, actionType int32) bool {
	toPublishMicroService := micro.NewService(micro.Registry(etcdInit.EtcdReg))
	toPublishService := from_user_proto.NewToPublishService("rpcUserService", toPublishMicroService.Client())
	var req from_user_proto.UpdateWorkCountRequest
	req.UserId = uid
	req.Count = count
	req.Type = actionType
	resp, err := toPublishService.UpdateWorkCount(context.TODO(), &req)
	if err != nil || resp.StatusCode != 0 {
		fmt.Println("work_count维护失败:", err)
		return false
	}
	return true

}
