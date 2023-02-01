package core

import (
	"context"
	"fmt"
	"time"
	"user/model"
	"user/rpc"
	"user/services"
	"user/utils"
)

/**
用户登录service层
*/
func (*UserService) Login(ctx context.Context, req *services.DouyinUserLoginRequest, resp *services.DouyinUserLoginResponse) error {
	fmt.Println("进入登录")
	username := req.Username
	password := req.Password
	if username == "" || password == "" {
		resp.StatusCode = -1
		resp.StatusMsg = "用户名或密码不能为空"
		resp.UserId = -1
		resp.Token = ""
		return nil
	}

	user, err := model.NewUserDaoInstance().FindUserByName(username)
	if err != nil {
		panic("调用UserDao的FindUserByName方法，根据用户名查询User失败")
		return err
	}

	if utils.Sha256(password) != user.Password {
		fmt.Println("用户名或密码错误")
		resp.StatusCode = -1
		resp.StatusMsg = "用户名或密码错误"
		resp.UserId = -1
		resp.Token = ""
		return nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = "登录成功"
	resp.UserId = user.UserId
	resp.Token = ""
	return nil
}

/**
注册 service层
*/
func (*UserService) Register(ctx context.Context, req *services.DouyinUserRegisterRequest, resp *services.DouyinUserRegisterResponse) error {
	fmt.Println("进入注册")
	//查询用户名，没有错误（能查到）
	username := req.Username
	password := req.Password

	if username == "" || password == "" {
		resp.StatusCode = -1
		resp.StatusMsg = "用户名或密码不能为空"
		resp.UserId = -1
		resp.Token = ""
		return nil
	}

	if _, err := model.NewUserDaoInstance().FindUserByName(username); err == nil {
		fmt.Println("用户名已经存在")
		resp.StatusCode = -1
		resp.StatusMsg = "用户名已经存在"
		resp.UserId = -1
		resp.Token = ""
		return nil
	}

	user := &model.User{
		Name:           username,
		Password:       utils.Sha256(password),
		FollowingCount: 0,
		FollowerCount:  0,
		CreateAt:       time.Now(),
	}

	_, err := model.NewUserDaoInstance().CreateUser(user)
	if err != nil {
		fmt.Println("创建用户失败")
		resp.StatusCode = -1
		resp.StatusMsg = "创建用户失败"
		resp.UserId = -1
		resp.Token = ""
		return err
	}

	user, _ = model.NewUserDaoInstance().FindUserByName(username)

	resp.StatusCode = 0
	resp.StatusMsg = "注册成功"
	resp.UserId = user.UserId
	resp.Token = ""
	return nil
}

/**
登录用户的详细信息 service层
*/
func (*UserService) UserInfo(ctx context.Context, req *services.DouyinUserRequest, resp *services.DouyinUserResponse) error {
	tokenUserIdConv := rpc.GetIdByToken(req.Token)
	fmt.Println(tokenUserIdConv)
	userId := req.UserId //传入的参数
	//tokenUserId := req.Token //token解析出来的userId
	//
	//if userId <= 0 || tokenUserId == "" {
	//	resp.StatusCode = -1
	//	resp.StatusMsg = "传入参数不全，不能为空"
	//	resp.User = &services.User{}
	//	return nil
	//}

	//curUserId, err := strconv.ParseInt(currentUserId, 10, 64) //当前用户的id
	//tokenUserIdConv, err := strconv.ParseInt(tokenUserId, 10, 64) //当前用户的id
	//if err != nil {
	//	resp.StatusCode = 1
	//	resp.StatusMsg = "类型转换失败"
	//	resp.User = &services.User{}
	//	return err
	//}

	//2. 根据userId查询User
	user, err := model.NewUserDaoInstance().FindUserById(userId)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "查找用户信息时发生异常"
		resp.User = &services.User{}
		return err
	}

	isFollow := false
	//这里可以做成根据tokenUserIdConv和userId查找relation表，判断isFollow tokenUserIdConv代表了当前用户的id,userId代表了想查找的id
	//如果当前用户查找的是自己的信息，那么是否关注的关系返回
	////TODO 这里应该调用Relation的微服务，是否有关注关系？为了不影响后续使用，目前先做了数据库查询，需要替换
	if _, err := model.NewUserDaoInstance().FindRelationById(userId, tokenUserIdConv); err == nil {
		//当前用户关注了user_id用户
		isFollow = true
	}

	resp.StatusCode = 0
	resp.StatusMsg = "查询用户信息成功"
	resp.User = BuildProtoUser(user, isFollow)
	return nil
}

func BuildProtoUser(item *model.User, isFollow bool) *services.User {
	user := services.User{
		Id:            item.UserId,
		Name:          item.Name,
		FollowCount:   item.FollowingCount,
		FollowerCount: item.FollowerCount,
		IsFollow:      isFollow,
	}
	return &user
}
