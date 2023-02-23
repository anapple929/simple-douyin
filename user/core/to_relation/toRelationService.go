package to_relation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"user/model"
	"user/rpc_server"
	to_user "user/services/from_relation"
	proto "user/services/to_relation"
	"user/utils/redis"
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
根据userId列表，获取User列表.
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

	var searchIds []int64        //保存从redis没命中的userId,统一去mysql中查一次
	var usersRedis []*model.User //从redis查到的user实体

	//把能从redis中查到的userId先都给查到，查不到的放到searchIds中，在数据库中统一给查出来
	for _, userIdSearchRedis := range req.UserId {
		count, err := redis.RdbUserId.Exists(redis.Ctx, strconv.FormatInt(userIdSearchRedis, 10)).Result()
		if err != nil {
			log.Println(err)
		}

		///***********************************///
		count = 0
		///***********************************///
		if count > 0 {
			var user *model.User //一个临时变量，存解构后的user实体
			//缓存里有
			//redis，先从redis通过userId查询user实体
			userString, err := redis.RdbUserId.Get(redis.Ctx, strconv.FormatInt(userIdSearchRedis, 10)).Result()
			if err != nil { //若查询缓存出错，则打印log
				//return 0, err
				log.Println("调用redis查询userId对应的信息出错", err)
			}
			json.Unmarshal([]byte(userString), &user)
			usersRedis = append(usersRedis, user)
		} else {
			//没查到，进入searchIds数组
			searchIds = append(searchIds, userIdSearchRedis)
		}
	}

	//调用数据库查user实体列表
	users, _ := model.NewUserDaoInstance().GetUsersByIds(searchIds)

	//把mysql查到的users数据存到redis中
	for _, user := range users {
		userValue, _ := json.Marshal(&user)
		_ = redis.RdbUserId.Set(redis.Ctx, strconv.FormatInt(user.UserId, 10), userValue, 0).Err()
	}

	//合并从mysql查到的users和redis查到的usersRedis
	users = append(users, usersRedis...)
	fmt.Println("得到的users是这样的")
	fmt.Println(users)

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
				Id:             user.UserId,
				Name:           user.Name,
				FollowCount:    user.FollowingCount,
				FollowerCount:  user.FollowerCount,
				IsFollow:       false,
				TotalFavorited: user.TotalFavorited,
				WorkCount:      user.WorkCount,
				FavoriteCount:  user.FavoriteCount,
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
			Id:             user.UserId,
			Name:           user.Name,
			FollowCount:    user.FollowingCount,
			FollowerCount:  user.FollowerCount,
			IsFollow:       isFollow,
			TotalFavorited: user.TotalFavorited,
			WorkCount:      user.WorkCount,
			FavoriteCount:  user.FavoriteCount,
		})
	}
	return userResult
}
