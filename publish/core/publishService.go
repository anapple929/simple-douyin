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

type PublishService struct {
}

//发视频
func (*PublishService) Publish(ctx context.Context, req *services.DouyinPublishActionRequest, resp *services.DouyinPublishActionResponse) error {
	//解析token,tokenUserIdConv是解析出来的上传视频的用户id
	tokenUserIdConv, err := rpc_server.GetIdByToken(req.Token)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "登录失效，请重新登录"
		return nil
	}

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
		resp.StatusCode = -1
		resp.StatusMsg = "创建视频失败"
		return err
	}

	resp.StatusCode = 0
	resp.StatusMsg = "上传视频成功"
	return nil
}

//视频列表
func (*PublishService) PublishList(ctx context.Context, req *services.DouyinPublishListRequest, resp *services.DouyinPublishListResponse) error {
	//解析token
	if req.Token == "" {
		resp.StatusCode = -1
		resp.StatusMsg = "登录失效，请重新登录"
		return nil
	}
	_, err := rpc_server.GetIdByToken(req.Token)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "登录失效，请重新登录"
		return nil
	}

	var videoResult []*services.Video

	//调用数据库查视频
	videos, err := model.NewVideoDaoInstance().QueryVideoByUserId(req.UserId)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "在video表查视频失败了"
		return err
	}

	//遍历实体Video，封装到VideoResult中
	for _, video := range videos {
		videoResult = append(videoResult, BuildProtoVideo(video, req.Token))
	}
	resp.StatusCode = 0
	resp.StatusMsg = "查询用户发布视频成功"
	resp.VideoList = videoResult

	fmt.Println(videoResult)
	return nil
}

func BuildProtoVideo(item *model.Video, token string) *services.Video {
	isFavorite := false
	userId, err := rpc_server.GetIdByToken(token)
	//没有错误，说明token存在且有效，userId是解析出的当前用户id
	if err != nil {
		userId = 0
	} else {
		isFavorite, err = rpc_server.GetFavoriteStatus(item.VideoId, userId)
		if err != nil {
			fmt.Println("调用远程favorite服务失败,错误原因是：")
			fmt.Println(err)
			return &services.Video{}
		}
	}

	video := services.Video{
		Id:            item.VideoId,
		Author:        BuildProtoUser(item.UserId, token),
		PlayUrl:       item.PlayUrl,
		CoverUrl:      item.CoverUrl,
		FavoriteCount: item.FavoriteCount,
		CommentCount:  item.CommentCount,
		IsFavorite:    isFavorite, //
		Title:         item.Title,
	}
	return &video
}

func BuildProtoUser(item_id int64, token string) *services.User {
	//根据id查user，封装成user
	rpcUserInfo, err := rpc_server.GetUserInfo(item_id, token)
	if err != nil {
		fmt.Println("调用远程user服务出错了，错误是：")
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
