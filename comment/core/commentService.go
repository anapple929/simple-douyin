package core

import (
	"comment/model"
	"comment/rpc_server"
	proto "comment/service"
	usersproto "comment/service/to_relation"
	userproto "comment/service/userproto"
	"comment/utils/redis"
	"context"
	"fmt"
	"log"
	"strconv"

	//"github.com/dtm-labs/client/dtmgrpc"
	//"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

type CommentService struct {
}

/*
*
评论
*/
func (*CommentService) CommentAction(ctx context.Context, in *proto.DouyinCommentActionRequest, out *proto.DouyinCommentActionResponse) error {
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

	//创建comment实体
	comment := &model.Comment{
		UserId:   userId,
		Content:  in.CommentText,
		VideoId:  in.VideoId,
		CreateAt: time.Now(),
	}

	//判断一下actionType的类型 1 发布消息 2 删除消息， 做一个if判断
	if in.ActionType == 1 {

		//barrier, _ := dtmgrpc.BarrierFromGrpc(ctx)
		//// 开启子事务屏障
		//db, _ := sqlx.NewMysql("").RawDB()
		//if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		//	//如果是发布消息，将拿到的参数调用model中的数据库方法，将数据传入数据库
		//	comment, _ = model.NewCommentDaoInstance().CreateComment(comment)
		//	//调用rpc_server的CommentCountAction，增加发布数
		//	rpc_server.CountAction(in.VideoId, 1, in.ActionType)
		//	return nil
		//}); err != nil {
		//	fmt.Println("全局事务出错了。")
		//}

		//没有分布式事务版本↓

		//如果是发布消息，将拿到的参数调用model中的数据库方法，将数据传入数据库
		comment, _ := model.NewCommentDaoInstance().CreateComment(comment)
		//调用rpc_server的CommentCountAction，增加发布数
		rpc_server.CountAction(in.VideoId, 1, in.ActionType)

		out.StatusMsg = "评论成功"
		out.StatusCode = 0
		out.Comment = ChangeComment(comment, user)
		return nil

	} else if in.ActionType == 2 {
		//如果是删除消息，将拿到的参数调用model中的数据库方法，删除记录
		//先判断一下删评论的人是不是删的自己的评论，如果不是，直接失败
		//通过commentId查询userId
		commentUserId, _ := model.NewCommentDaoInstance().GetUserIdByCommentId(in.CommentId)
		if userId != commentUserId {
			out.StatusMsg = "不能删除别人的评论"
			out.StatusCode = 0
			out.Comment = ChangeComment(comment, user)
			return nil
		}
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
	//评论状态改变，删除缓存
	key := strconv.FormatInt(in.VideoId, 10)
	countRedis, err := redis.RdbVideoId.Exists(redis.Ctx, key).Result()
	if err != nil {
		log.Println(err)
	}
	if countRedis > 0 {
		fmt.Println("删除了video key存的缓存")
		redis.RdbVideoId.Del(redis.Ctx, key)
	}

	return nil
}

/*
*
评论列表
*/
func (*CommentService) CommentList(ctx context.Context, in *proto.DouyinCommentListRequest, out *proto.DouyinCommentListResponse) error {
	//存评论列表
	var commentResult []*proto.Comment
	//拿到userIds集合，调用usersinfo方法，查一批User实体
	var userIds []int64
	//调用数据库方法
	comments, _ := model.NewCommentDaoInstance().QueryComment(in.VideoId)

	for _, comment := range comments {
		userIds = append(userIds, comment.UserId)
	}
	fmt.Println(userIds)
	//调用usersinfo方法，查一批User实体
	users, _ := rpc_server.GetUsersInfo(userIds, "")
	fmt.Println(users)
	for _, comment := range comments {
		for _, user := range users {
			if user.Id == comment.UserId {
				commentResult = append(commentResult, BuildProtoComment(comment, user))
				break
			}
		}
	}

	out.CommentList = commentResult
	out.StatusCode = 0
	out.StatusMsg = "查询成功"
	fmt.Println(commentResult)
	return nil
}

func BuildProtoComment(comment *model.Comment, user *usersproto.User) *proto.Comment {
	return &proto.Comment{
		Id:         comment.CommentId,
		User:       BuildProtoUser(user),
		Content:    comment.Content,
		CreateDate: comment.CreateAt.Format("2006-01-02 15:04:05"),
	}
}

func BuildProtoUser(user *usersproto.User) *proto.User {
	return &proto.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
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
		CreateDate: comment.CreateAt.Format("2006-01-02 15:04:05"),
	}
}
