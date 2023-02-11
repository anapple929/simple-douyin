package tofavorite

import (
	"context"
	"errors"
	"publish/model"
	"publish/rpc_server"
	"publish/services/favorite_to_video_proto"
	proto "publish/services/to_favorite"
	usersproto "publish/services/to_relation"
)

type ToFavoriteService struct {
}

/**
给Favorite微服务调用，更新视频表的点赞数。
req携带的参数：videoId 视频id   count 增加或者减少的数字   type 1增加2减少
*/
func (ToFavoriteService) UpdateFavoriteCount(ctx context.Context, req *proto.UpdateFavoriteCountRequest, resp *proto.UpdateFavoriteCountResponse) error {
	if req.VideoId <= 0 || (req.Type != 1 && req.Type != 2) {
		resp.StatusCode = -1
		return errors.New("传入的videoId或者type有误")
	}
	//查一下，这个videoId能否查到，查不到报错，查到了返回count
	if _, err := model.NewVideoDaoInstance().FindVideoById(req.VideoId); err != nil {
		return errors.New("传入的VideoId查不到")
	}
	//调用数据库的修改功能
	if req.Type == 1 {
		//增加
		model.NewVideoDaoInstance().AddFavoriteCount(req.VideoId, req.Count)
	} else if req.Type == 2 {
		//减少
		model.NewVideoDaoInstance().ReduceFavoriteCount(req.VideoId, req.Count)
	}

	resp.StatusCode = 0
	return nil
}

/**
根据videoId列表，获取Video列表
*/
func (*ToFavoriteService) GetVideosByIds(ctx context.Context, req *proto.GetVideosByIdsRequest, resp *proto.GetVideosByIdsResponse) error {
	//解析token
	if req.Token == "" {
		resp.StatusCode = -1
		return nil
	}
	user_id, err := rpc_server.GetIdByToken(req.Token) //解析出的userId，用于找是否关注了视频，是否和一些用户有关注关系
	if err != nil {
		resp.StatusCode = -1
		return nil
	}

	//将video实体封装成resp.Video类型
	var videoResult []*proto.Video

	if len(req.VideoId) == 0 {
		resp.StatusCode = 0
		resp.VideoList = videoResult
		return nil
	}
	//调用数据库查video实体列表
	videos, err := model.NewVideoDaoInstance().GetVideosByIds(req.VideoId)
	if err != nil {
		resp.StatusCode = -1
		resp.VideoList = nil
		return errors.New("调用数据库出错！")
	}

	//for _, video := range videos {
	//	videoResult = append(videoResult, BuildProtoVideo(video, req.Token))
	//}

	//拿到userIds集合，调用usersinfo方法，查一批User实体
	var userIds []int64
	var favoriteStatus []*favorite_to_video_proto.FavoriteStatus
	for _, video := range videos {
		//封装userIds，作为参数，调用user微服务的远程接口
		userIds = append(userIds, video.UserId)
		//封装isFavorites，作为参数，调用favorite微服务的远程接口
		favoriteStatus = append(favoriteStatus, &favorite_to_video_proto.FavoriteStatus{UserId: user_id, VideoId: video.VideoId, IsFavorite: false})
	}
	//调用usersinfo方法，查一批User实体
	users, err := rpc_server.GetUsersInfo(userIds, req.Token)
	isFavorites, _ := rpc_server.GetFavoritesStatus(favoriteStatus)

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
	resp.VideoList = videoResult
	return nil
}

func BuildProtoVideo(video *model.Video, user *usersproto.User, isFavorite bool) *proto.Video {
	return &proto.Video{
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

func BuildProtoUser(user *usersproto.User) *proto.User {
	return &proto.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}

//
///**
//构造一个控制层Video对象
//*/
//func BuildProtoVideo(item *model.Video, token string) *proto.Video {
//	isFavorite := false
//	userId, err := rpc_server.GetIdByToken(token)
//	fmt.Println(userId)
//	//没有错误，说明token存在且有效，userId是解析出的当前用户id
//	if err == nil {
//		isFavorite, err = rpc_server.GetFavoriteStatus(item.VideoId, userId)
//		if err != nil {
//			fmt.Println("调用远程favorite服务失败,错误原因是：")
//			fmt.Println(err)
//			return &proto.Video{}
//		}
//	}
//
//	video := proto.Video{
//		Id:            item.VideoId,
//		Author:        BuildProtoUser(item.UserId, token),
//		PlayUrl:       item.PlayUrl,
//		CoverUrl:      item.CoverUrl,
//		FavoriteCount: item.FavoriteCount,
//		CommentCount:  item.CommentCount,
//		IsFavorite:    isFavorite,
//		Title:         item.Title,
//	}
//
//	return &video
//}
//
///**
//构造一个控制层User对象
//*/
//func BuildProtoUser(item_id int64, token string) *proto.User {
//	rpcUserInfo, err := rpc_server.GetUserInfo(item_id, token)
//	if err != nil {
//		fmt.Println("调用远程user服务出错了，错误是：")
//		fmt.Println(err)
//		return &proto.User{}
//	}
//	//如果是空，没登陆，返回的应该是默认值
//	if rpcUserInfo == nil {
//		return &proto.User{}
//	}
//	user := proto.User{
//		Id:            rpcUserInfo.Id,
//		Name:          rpcUserInfo.Name,
//		FollowCount:   rpcUserInfo.FollowCount,
//		FollowerCount: rpcUserInfo.FollowerCount,
//		IsFollow:      rpcUserInfo.IsFollow,
//	}
//
//	return &user
//}
