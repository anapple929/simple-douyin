package core

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
	"user/model"
	"user/rpc_server"
	"user/services"
	"user/utils/redis"
	"user/utils/sha256"
)

type UserService struct {
}

/**
用户登录，service层，
req:用户名，密码
resp:
*/
func (*UserService) Login(ctx context.Context, req *services.DouyinUserLoginRequest, resp *services.DouyinUserLoginResponse) error {
	//获取输入的用户名和密码
	username := req.Username
	password := req.Password

	//判断用户名和密码是否为空
	if username == "" || password == "" {
		resp.StatusCode = -1
		resp.StatusMsg = "用户名或密码不能为空"
		resp.UserId = -1
		resp.Token = ""
		return nil
	}

	//调用数据库方法，查找是否有这个用户名的实体
	user, err := model.NewUserDaoInstance().FindUserByName(username)
	if err != nil {
		return err
	}

	//判断密码是否正确
	if sha256.Sha256(password) != user.Password {
		resp.StatusCode = -1
		resp.StatusMsg = "用户名或密码错误"
		resp.UserId = -1
		resp.Token = ""
		return nil
	}

	//将返回值设置在resp中，token在网关签发
	resp.StatusCode = 0
	resp.StatusMsg = "登录成功"
	resp.UserId = user.UserId
	return nil
}

/**
注册 service层
req:用户名，密码
resp:新生成的userId，
*/
//TODO 有查询是否有用户和插入数据，最好做个事务，防止有同名用户
func (*UserService) Register(ctx context.Context, req *services.DouyinUserRegisterRequest, resp *services.DouyinUserRegisterResponse) error {
	//在req中获取用户名和密码
	username := req.Username
	password := req.Password

	//用户名和密码为空，返回
	if username == "" || password == "" {
		resp.StatusCode = -1
		resp.StatusMsg = "用户名或密码不能为空"
		resp.UserId = -1
		resp.Token = ""
		return nil
	}

	//调用数据库方法，查询是否有同名实体
	if _, err := model.NewUserDaoInstance().FindUserByName(username); err == nil {
		resp.StatusCode = -1
		resp.StatusMsg = "用户名已经存在"
		resp.UserId = -1
		resp.Token = ""
		return nil
	}

	//创建一个dao层User实体
	user := &model.User{
		Name:           username,
		Password:       sha256.Sha256(password),
		FollowingCount: 0,
		FollowerCount:  0,
		CreateAt:       time.Now(),
	}

	//调用数据库方法，创建一个新的User实体
	_, err := model.NewUserDaoInstance().CreateUser(user)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "创建用户失败"
		resp.UserId = -1
		resp.Token = ""
		return err
	}

	//根据用户名，查询新用户的userId，作为返回值返回
	user, _ = model.NewUserDaoInstance().FindUserByName(username)

	//补充resp
	resp.StatusCode = 0
	resp.StatusMsg = "注册成功"
	resp.UserId = user.UserId
	resp.Token = ""
	return nil
}

/**
登录用户的详细信息 service层
在登录后或注册后被调用，查看自己的信息。
在刷视频的时候，点进作者头像调用，查看别人的信息。
resp: token判断用户是否登录， userId，想要查询的用户的Id
*/
//TODO，查出来的数据可以放在缓存里
func (*UserService) UserInfo(ctx context.Context, req *services.DouyinUserRequest, resp *services.DouyinUserResponse) error {
	fmt.Println("进入userInfo方法了")
	if req.Token == "" { //如果传进来的直接是空，比如feed流可以无登录用户刷信息，video会调用到这里，可以不用调用rpc的token去解析，直接返回。
		resp.StatusCode = -1
		resp.StatusMsg = "登录失效，请重新登录"
		resp.User = &services.User{}
		return nil
	}

	//解析token,从token解析出userId，能解析出的才能查询用户信息，否则要先登录
	tokenUserIdConv, err := rpc_server.GetIdByToken(req.Token)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "登录失效，请重新登录"
		resp.User = &services.User{}
		return nil
	}
	// 获得想要获取详细信息的userId
	userId := req.UserId

	var user *model.User
	var userString string

	count, err := redis.RdbUserId.Exists(redis.Ctx, strconv.FormatInt(userId, 10)).Result()
	if err != nil {
		log.Println(err)
	}

	if count > 0 {
		//缓存里有
		//redis，先从redis通过userId查询user实体
		userString, err = redis.RdbUserId.Get(redis.Ctx, strconv.FormatInt(userId, 10)).Result()
		if err != nil { //若查询缓存出错，则打印log
			//return 0, err
			log.Println("调用redis查询userId对应的信息出错", err)
		}
		json.Unmarshal([]byte(userString), &user)
		fmt.Println("redis查出来的结果")
		fmt.Println(&user)
	} else {
		fmt.Println("查数据库")
		//根据userId查询User
		user, err = model.NewUserDaoInstance().FindUserById(userId)
		fmt.Println("输出一下改了protobuf之后的user")
		fmt.Println(user)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = "查找用户信息时发生异常"
			return err
		}
		//把查到的数据放入redis
		userValue, _ := json.Marshal(&user)
		_ = redis.RdbUserId.Set(redis.Ctx, strconv.FormatInt(userId, 10), userValue, 0).Err()
	}

	fmt.Println(user)

	//TODO 这里应该调用Relation的微服务，是否有关注关系？为了不影响后续使用，目前先做了数据库查询，需要替换
	isFollow, err := model.NewUserDaoInstance().FindRelationById(userId, tokenUserIdConv)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "查询relation数据库，两人是否有关注关系的时候失败"
		return err
	}

	resp.StatusCode = 0
	resp.StatusMsg = "查询用户信息成功"
	resp.User = BuildProtoUser(user, isFollow)
	return nil
}

/**
构建视图层User
*/
func BuildProtoUser(item *model.User, isFollow bool) *services.User {
	user := services.User{
		Id:             item.UserId,
		Name:           item.Name,
		FollowCount:    item.FollowingCount,
		FollowerCount:  item.FollowerCount,
		IsFollow:       isFollow,
		TotalFavorited: item.TotalFavorited,
		WorkCount:      item.WorkCount,
		FavoriteCount:  item.FavoriteCount,
	}
	return &user
}
