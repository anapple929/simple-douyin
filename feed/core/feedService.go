package core

import (
	"context"
	"feed/model"
	"feed/rpc"
	"feed/services"
	"fmt"
	"time"
)

func (*FeedService) Feed(ctx context.Context, req *services.DouyinFeedRequest, resp *services.DouyinFeedResponse) error {
	fmt.Println("feedList service层")

	fmt.Println("拿到的参数")
	fmt.Println(req.Token)
	fmt.Println(req.LatestTime)

	LatestTime := req.LatestTime
	format := "2006-01-02 15:04:05"
	t := time.Unix(LatestTime/1000, 0)
	searchTime := t.Format(format)
	//日期转换
	//formatTime, _ := time.Parse("2006-01-02 15:04:05", searchTime)
	var videoResult []*services.Video

	videos := model.NewVideoDaoInstance().QueryVideo(&searchTime, 30)

	//遍历实体Video，封装到VideoResult中
	for _, video := range videos {
		videoResult = append(videoResult, BuildProtoVideo(video))
	}
	resp.StatusCode = 0
	resp.StatusMsg = "查询视频成功"
	resp.VideoList = videoResult

	fmt.Println(videoResult)
	return nil
}

func BuildProtoVideo(item *model.Video) *services.Video {
	video := services.Video{
		Id:            item.VideoId,
		Author:        BuildProtoUser(item.UserId),
		PlayUrl:       item.PlayUrl,
		CoverUrl:      item.CoverUrl,
		FavoriteCount: item.FavoriteCount,
		CommentCount:  item.CommentCount,
		IsFavorite:    false, // //TODO 这里需要调用点赞表，来获取一个用户是否喜欢这个视频，目前传成了false
		Title:         item.Title,
	}
	return &video
}

func BuildProtoUser(item_id int64) *services.User {
	//根据id查user，封装成user //TODO，调用，用户的tokenrpc还没完全封装好，明天封装token(已完成)
	rpcUserInfo := rpc.GetUserInfo(item_id, "")
	user := services.User{
		Id:            rpcUserInfo.Id,
		Name:          rpcUserInfo.Name,
		FollowCount:   rpcUserInfo.FollowCount,
		FollowerCount: rpcUserInfo.FollowerCount,
		IsFollow:      rpcUserInfo.IsFollow,
	}
	return &user
}
