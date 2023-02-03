package core

import (
	"context"
	"fmt"
	"publish/model"
	"publish/rpc_server"
	"publish/services"
	utils "publish/utils"
	"time"
)

var token string

func (*PublishService) Publish(ctx context.Context, req *services.DouyinPublishActionRequest, resp *services.DouyinPublishActionResponse) error {
	fmt.Println("publish service层")

	//获取userId
	tokenUserIdConv, _ := rpc_server.GetIdByToken(req.Token)
	//tokenUserIdConv, err := strconv.ParseInt(tokenUserId, 10, 64) //token解析出来的userId
	//
	// = rpc_server.GetIdByToken(req.Token)

	//if tokenUserId == "" {
	//	resp.StatusCode = -1
	//	resp.StatusMsg = "token失效，请先登录后操作"
	//	return nil
	//}
	//获取title
	title := req.Title
	//上传视频
	videoUrl := utils.UploadVideo(req.Data)
	fmt.Println("上传视频地址是" + videoUrl)

	//获取封面
	coverBytes, _ := utils.ReadFrameAsJpeg(videoUrl)
	//上传封面
	coverUrl := utils.UploadPicture(coverBytes)
	fmt.Println("上传视频封面的地址是" + coverUrl)

	//构造video Dao模型
	video := &model.Video{
		UserId:        tokenUserIdConv,
		Title:         title,
		CoverUrl:      coverUrl,
		PlayUrl:       videoUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		CreateAt:      time.Now(),
	}

	//调用数据库创建Video方法
	if _, err := model.NewVideoDaoInstance().CreateVideo(video); err != nil {
		fmt.Println("创建视频失败")
		resp.StatusCode = -1
		resp.StatusMsg = "创建视频失败"
		return err
	}

	fmt.Println("到这里了，返回")
	resp.StatusCode = 0
	resp.StatusMsg = "上传视频成功"
	return nil
}

func (*PublishService) PublishList(ctx context.Context, req *services.DouyinPublishListRequest, resp *services.DouyinPublishListResponse) error {
	fmt.Println("publishList service层")
	token = req.Token

	var videoResult []*services.Video

	videos, _ := model.NewVideoDaoInstance().QueryVideoByUserId(req.UserId)

	//遍历实体Video，封装到VideoResult中
	for _, video := range videos {
		videoResult = append(videoResult, BuildProtoVideo(video))
	}
	resp.StatusCode = 0
	resp.StatusMsg = "查询用户发布视频成功"
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
		IsFavorite:    rpc_server.GetFavoriteStatus(item.VideoId, item.UserId), //
		Title:         item.Title,
	}
	return &video
}

func BuildProtoUser(item_id int64) *services.User {
	//根据id查user，封装成user
	rpcUserInfo, _ := rpc_server.GetUserInfo(item_id, token)
	user := services.User{
		Id:            rpcUserInfo.Id,
		Name:          rpcUserInfo.Name,
		FollowCount:   rpcUserInfo.FollowCount,
		FollowerCount: rpcUserInfo.FollowerCount,
		IsFollow:      rpcUserInfo.IsFollow,
	}
	return &user
}
