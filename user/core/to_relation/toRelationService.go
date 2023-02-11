package to_relation

import (
	"context"
	"errors"
	"fmt"
	"user/model"
	"user/rpc_server"
	to_user "user/services/from_relation"
	proto "user/services/to_relation"
)

type ToRelationService struct {
}

/**
给Relation微服务调用，更新用户表的follower_count。
req携带的参数：userId 用户id   count 增加或者减少的数字   type 1增加2减少
*/
func (ToRelationService) UpdateFollowerCount(ctx context.Context, req *proto.UpdateFollowerCountRequest, resp *proto.UpdateFollowerCountResponse) error {
	if req.UserId <= 0 || (req.Type != 1 && req.Type != 2) {
		resp.StatusCode = -1
		return errors.New("传入的userId或者type有误")
	}
	//查一下，这个videoId能否查到，查不到报错，查到了返回count
	if _, err := model.NewUserDaoInstance().FindUserById(req.UserId); err != nil {
		return errors.New("传入的VideoId查不到")
	}
	//调用数据库的修改功能
	if req.Type == 1 {
		//增加
		model.NewUserDaoInstance().AddFollowerCount(req.UserId, req.Count)
	} else if req.Type == 2 {
		//减少
		model.NewUserDaoInstance().ReduceFollowerCount(req.UserId, req.Count)
	}

	resp.StatusCode = 0
	return nil
}

/**
给Relation微服务调用，更新用户表的following_count。
req携带的参数：userId 用户id   count 增加或者减少的数字   type 1增加2减少
*/
func (ToRelationService) UpdateFollowingCount(ctx context.Context, req *proto.UpdateFollowingCountRequest, resp *proto.UpdateFollowingCountResponse) error {
	if req.UserId <= 0 || (req.Type != 1 && req.Type != 2) {
		resp.StatusCode = -1
		return errors.New("传入的userId或者type有误")
	}
	//查一下，这个videoId能否查到，查不到报错，查到了返回count
	if _, err := model.NewUserDaoInstance().FindUserById(req.UserId); err != nil {
		return errors.New("传入的VideoId查不到")
	}
	//调用数据库的修改功能
	if req.Type == 1 {
		//增加
		model.NewUserDaoInstance().AddFollowingCount(req.UserId, req.Count)
	} else if req.Type == 2 {
		//减少
		model.NewUserDaoInstance().ReduceFollowingCount(req.UserId, req.Count)
	}

	resp.StatusCode = 0
	return nil
}

/**
根据userId列表，获取User列表
*/
func (ToRelationService) GetUsersByIds(ctx context.Context, req *proto.GetUsersByIdsRequest, resp *proto.GetUsersByIdsResponse) error {
	var tokenUserIdConv int64
	tokenUserIdConv = -1
	//解析token,从token解析出userId，能解析出的才能查询用户信息，否则要先登录
	if req.Token != "" {
		tokenUserIdConv, _ = rpc_server.GetIdByToken(req.Token)
	}

	//将User实体封装成resp.User类型
	var userResult []*proto.User

	if len(req.UserId) == 0 {
		resp.StatusCode = 0
		resp.UserList = userResult
		return nil
	}
	fmt.Println("用户id列表是：")
	fmt.Println(req.UserId)
	//调用数据库查user实体列表
	users, _ := model.NewUserDaoInstance().GetUsersByIds(req.UserId)

	if tokenUserIdConv != -1 { //有登录的token，并且解析出来了，才远程调用两个人是否有关系的函数。
		//构造很多RelationStatus结构体，形成一个结构体数组，传进去
		var relationsStatus []*to_user.RelationStatus
		for _, user := range users {
			relationsStatus = append(relationsStatus, &to_user.RelationStatus{FollowingId: tokenUserIdConv, FollowerId: user.UserId, IsFollow: false})
		}
		//TODO 远程调用
		relations, _ := model.NewRelationDaoInstance().GetRelationsByIds(relationsStatus) //远程调用应该再传一个token
		//通过relations和users build返回的user
		userResult = BuildProtoUser(users, relations)
	} else {
		for _, user := range users {
			userResult = append(userResult, &proto.User{
				Id:            user.UserId,
				Name:          user.Name,
				FollowCount:   user.FollowingCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      false,
			})
		}
	}

	fmt.Println(userResult)
	resp.StatusCode = 0
	resp.UserList = userResult
	return nil
}

/**
构造一个控制层User对象
*/
func BuildProtoUser(users []*model.User, relations []*to_user.RelationStatus) []*proto.User {
	var userResult []*proto.User
	for _, user := range users {
		isFollow := false
		for _, relation := range relations {
			if relation.FollowerId == user.UserId {
				isFollow = relation.IsFollow
				break
			}
		}
		userResult = append(userResult, &proto.User{
			Id:            user.UserId,
			Name:          user.Name,
			FollowCount:   user.FollowingCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      isFollow,
		})
	}
	return userResult
}
