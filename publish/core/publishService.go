package core

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"publish/model"
	"publish/rpc_server"
	"publish/services"
	"publish/services/favorite_to_video_proto"
	usersproto "publish/services/to_relation"
	utils "publish/utils"
	"sync"
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
	//生成视频地址
	videoUUID, _ := uuid.NewV4()
	videoDir := time.Now().Format("2006-01-02") + "/" + videoUUID.String() + ".mp4"
	videoUrl := "https://" + "simple-douyin-1122233" + ".oss-cn-hangzhou.aliyuncs.com/" + videoDir
	fmt.Println("上传视频地址是" + videoUrl)
	//生成图片地址
	pictureUUID, _ := uuid.NewV4()
	pictureDir := time.Now().Format("2006-01-02") + "/" + pictureUUID.String() + ".jpg"
	coverUrl := "https://" + "simple-douyin-1122233" + ".oss-cn-hangzhou.aliyuncs.com/" + pictureDir
	fmt.Println("上传视频封面的地址是" + coverUrl)

	//开启协程上传
	go func() {
		//上传视频
		_ = utils.UploadVideo(videoDir, req.Data)
		//time.Sleep(2*time.Second)
		//获取封面
		coverBytes, _ := utils.ReadFrameAsJpeg(videoUrl)
		//上传封面
		_ = utils.UploadPicture(pictureDir, coverBytes)
	}()
	rpc_server.UpdateWorkCount(tokenUserIdConv, 1, 1) //调用远程user方法，更新work_count字段，+1

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
	user_id, err := rpc_server.GetIdByToken(req.Token)
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

	//拿到userIds集合，调用usersinfo方法，查一批User实体
	var userIds []int64
	var favoriteStatus []*favorite_to_video_proto.FavoriteStatus
	for _, video := range videos {
		//封装userIds，作为参数，调用user微服务的远程接口
		userIds = append(userIds, video.UserId)
		//封装isFavorites，作为参数，调用favorite微服务的远程接口
		favoriteStatus = append(favoriteStatus, &favorite_to_video_proto.FavoriteStatus{UserId: user_id, VideoId: video.VideoId, IsFavorite: false})
	}

	//用协程去调用两个微服务，批量查询user实体和favoriteStatus实体
	var users []*usersproto.User
	var isFavorites []*favorite_to_video_proto.FavoriteStatus

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		//调用usersinfo方法，查一批User实体
		users, err = rpc_server.GetUsersInfo(userIds, req.Token)
	}()

	go func() {
		defer wg.Done()
		//调用FavoritesStatus方法，查一批FavoriteStatus实体
		isFavorites, _ = rpc_server.GetFavoritesStatus(favoriteStatus)
	}()
	wg.Wait()

	//如果查到的users的某一项id和video的id的某一项一致，那么就把user封装到返回的video中。
	//如果查到的isfavorite的某一项id和video的id的某一项一致，那么就把user封装到返回的video中。
	for _, video := range videos {
		isFavorite := false
		//遍历isFavorites，找userid和videoid对应的点赞状态
		for _, isF := range isFavorites {
			if isF.UserId == user_id && isF.VideoId == video.VideoId { //当前用户，找到了videoId对应的点赞状态
				isFavorite = isF.IsFavorite
				break
			}
		}
		//遍历users，根据video中的userId找user实体
		for _, user := range users {
			if video.UserId == user.Id { //videoId找到它对应的User实体了
				videoResult = append(videoResult, BuildProtoVideo(video, user, isFavorite))
				break
			}
		}
	}

	resp.StatusCode = 0
	resp.StatusMsg = "查询用户发布视频成功"
	resp.VideoList = videoResult

	fmt.Println(videoResult)
	return nil
}

func BuildProtoVideo(video *model.Video, user *usersproto.User, isFavorite bool) *services.Video {
	return &services.Video{
		Id:            video.VideoId,
		Author:        BuildProtoUser(user),
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    isFavorite,
		Title:         video.Title,
	}
}

func BuildProtoUser(user *usersproto.User) *services.User {
	return &services.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}
