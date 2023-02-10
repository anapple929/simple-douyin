package core

import (
	"comment/model"
	"comment/rpc_server"
	proto "comment/service"
	userproto "comment/service/userproto"
	"context"
	"fmt"
)

type CommentService struct {
}

/**
评论
*/
func (*CommentService) CommentAction(ctx context.Context, in *proto.DouyinCommentActionRequest, out *proto.DouyinCommentActionResponse) error {
	fmt.Println("comment service层 commentAction")
	fmt.Println("拿到的参数是")
	fmt.Println(in.CommentId)
	fmt.Println(in.ActionType)
	fmt.Println(in.Token)
	fmt.Println(in.CommentText)
	fmt.Println(in.VideoId)
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

	//判断一下actionType的类型 1 发布消息 2 删除消息， 做一个if判断
	if in.ActionType == 1 {
		//如果是发布消息，将拿到的参数调用model中的数据库方法，将数据传入数据库
		comment, _ := model.NewCommentDaoInstance().CreateComment(&model.Comment{})
		//调用rpc_server的CommentCountAction，增加发布数
		rpc_server.CountAction(in.VideoId, 1, in.ActionType)

		out.StatusMsg = "评论成功"
		out.StatusCode = 0
		out.Comment = ChangeComment(comment, user)
		fmt.Println(out.Comment)

		return nil

	} else if in.ActionType == 2 {
		//如果是删除消息，将拿到的参数调用model中的数据库方法，删除记录
		err := model.NewCommentDaoInstance().DeleteCommentById(in.CommentId)
		if err != nil {
			out.StatusCode = -1
			out.StatusMsg = "调用数据库删除方法时出现异常"
			return nil
		}
		//调用rpc_server的CommentCountAction，减少发布数
		rpc_server.CountAction(in.VideoId, 1, in.ActionType)
		out.StatusMsg = "删除评论成功"
		out.StatusCode = 0
		out.Comment = &proto.Comment{}
	} else {
		out.StatusCode = -1
		out.StatusMsg = "actionType有问题"
		return nil
	}

	return nil
}

/**
评论列表
*/
func (*CommentService) CommentList(ctx context.Context, in *proto.DouyinCommentListRequest, out *proto.DouyinCommentListResponse) error {
	fmt.Println("comment service层 commentList")

	var commentResult []*proto.Comment
	comments := model.NewCommentDaoInstance().QueryComment(in.VideoId)
	for _, comment := range comments {
		commentResult = append(commentResult, BuildProtoComment(comment, in.Token))
	}

	out.CommentList = commentResult
	out.StatusCode = 0
	out.StatusMsg = "查询成功"
	fmt.Println(commentResult)
	return nil
}

func BuildProtoComment(comment *model.Comment, token string) *proto.Comment {
	return &proto.Comment{
		Id:         comment.CommentId,
		User:       BuildProtoUser(comment.UserId, token),
		Content:    comment.Content,
		CreateDate: comment.CreateAt.String(),
	}
}

func BuildProtoUser(item_id int64, token string) *proto.User {
	rpcUserInfo, err := rpc_server.GetUserInfo(item_id, token)
	if err != nil {
		fmt.Println("调用远程user服务失败,错误原因是：")
		fmt.Println(err)
		return &proto.User{}
	}
	//如果是空，没登陆，返回的应该是默认值
	if rpcUserInfo == nil {
		return &proto.User{}
	}
	user := proto.User{
		Id:            rpcUserInfo.Id,
		Name:          rpcUserInfo.Name,
		FollowCount:   rpcUserInfo.FollowCount,
		FollowerCount: rpcUserInfo.FollowerCount,
		IsFollow:      rpcUserInfo.IsFollow,
	}
	return &user
}

func ChangeUser(user *userproto.User) *proto.User {
	return &proto.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}
func ChangeComment(comment *model.Comment, user *userproto.User) *proto.Comment {
	return &proto.Comment{
		Id:         comment.CommentId,
		User:       ChangeUser(user),
		Content:    comment.Content,
		CreateDate: comment.CreateAt.String(),
	}
}
