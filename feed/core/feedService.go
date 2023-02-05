package core

import (
	"context"
	"feed/model"
	"feed/rpc_server"
	"feed/services"
	"fmt"
	"time"
)

type FeedService struct {
}

//数据流
func (*FeedService) Feed(ctx context.Context, req *services.DouyinFeedRequest, resp *services.DouyinFeedResponse) error {
	LatestTime := req.LatestTime
	//日期转换
	format := "2006-01-02 15:04:05"
	t := time.Unix(LatestTime/1000, 0)
	searchTime := t.Format(format)

	//调用数据库方法，查询视频
	videos := model.NewVideoDaoInstance().QueryVideo(&searchTime, 5)

	//遍历实体Video，封装到VideoResult中
	var videoResult []*services.Video
	for _, video := range videos {
		videoResult = append(videoResult, BuildProtoVideo(video, req.Token))
	}
	//返回
	resp.StatusCode = 0
	resp.StatusMsg = "查询视频成功"
	resp.VideoList = videoResult

	return nil
}

func BuildProtoVideo(item *model.Video, token string) *services.Video {
	isFavorite := false
	var userId int64
	var err error
	if token != "" { //未登录用户
		userId, err = rpc_server.GetIdByToken(token)
		//没有错误，说明token存在且有效，userId是解析出的当前用户id
		if err == nil {
			//登录状态异常，也可以刷视频，用户信息查不出来
			isFavorite, err = rpc_server.GetFavoriteStatus(item.VideoId, userId)
			if err != nil {
				fmt.Println("调用远程favorite服务失败,错误原因是：")
				fmt.Println(err)
				return &services.Video{}
			}
		}
	} else {
		userId = 0
	}

	video := services.Video{
		Id:            item.VideoId,
		Author:        BuildProtoUser(item.UserId, token),
		PlayUrl:       item.PlayUrl,
		CoverUrl:      item.CoverUrl,
		FavoriteCount: item.FavoriteCount,
		CommentCount:  item.CommentCount,
		IsFavorite:    isFavorite,
		Title:         item.Title,
	}
	return &video
}

func BuildProtoUser(item_id int64, token string) *services.User {
	rpcUserInfo, err := rpc_server.GetUserInfo(item_id, token)
	if err != nil {
		fmt.Println("调用远程user服务失败,错误原因是：")
		fmt.Println(err)
		return &services.User{}
	}
	//如果是空，没登陆，返回的应该是默认值
	if rpcUserInfo == nil {
		return &services.User{}
	}
	user := services.User{
		Id:            rpcUserInfo.Id,
		Name:          rpcUserInfo.Name,
		FollowCount:   rpcUserInfo.FollowCount,
		FollowerCount: rpcUserInfo.FollowerCount,
		IsFollow:      rpcUserInfo.IsFollow,
	}
	return &user
}
