package core

import (
	"context"
	"feed/model"
	"feed/rpc_server"
	"feed/services"
	"feed/services/favorite_to_video_proto"
	usersproto "feed/services/to_relation"
	"sync"
	"time"
)

type FeedService struct {
}

//数据流
func (*FeedService) Feed(ctx context.Context, req *services.DouyinFeedRequest, resp *services.DouyinFeedResponse) error {
	var user_id int64
	user_id = -1
	if req.Token != "" {
		user_id, _ = rpc_server.GetIdByToken(req.Token)
	}

	LatestTime := req.LatestTime
	//日期转换
	format := "2006-01-02 15:04:05"
	t := time.Unix(LatestTime/1000, 0)
	searchTime := t.Format(format)

	//调用数据库方法，查询视频
	videos := model.NewVideoDaoInstance().QueryVideo(&searchTime, 5)

	var videoResult []*services.Video

	//拿到userIds集合，调用usersinfo方法，查一批User实体
	var userIds []int64                                          //远程调用批量查询user实体的输入参数
	var favoriteStatus []*favorite_to_video_proto.FavoriteStatus //远程调用favorite批量查询是否点赞的输入参数
	var users []*usersproto.User                                 //接收远程调用user返回的user实体
	var isFavorites []*favorite_to_video_proto.FavoriteStatus    //接收远程调用Favorite返回的是否点赞实体

	if user_id == -1 { //没登陆，不用远程调用favorite了。
		for _, video := range videos {
			//封装userIds，作为参数，调用user微服务的远程接口
			userIds = append(userIds, video.UserId)
		}
		users, _ = rpc_server.GetUsersInfo(userIds, req.Token)
		//如果查到的users的某一项id和video的id的某一项一致，那么就把user封装到返回的video中。
		for _, video := range videos {
			//遍历users，根据video中的userId找user实体
			for _, user := range users {
				if video.UserId == user.Id { //videoId找到它对应的User实体了
					videoResult = append(videoResult, BuildProtoVideo(video, user, false))
					break
				}
			}
			//videoResult = append(videoResult, BuildProtoVideo(video, &usersproto.User{}, false))
		}
	} else {
		for _, video := range videos {
			//封装userIds，作为参数，调用user微服务的远程接口
			userIds = append(userIds, video.UserId)
			//封装isFavorites，作为参数，调用favorite微服务的远程接口
			favoriteStatus = append(favoriteStatus, &favorite_to_video_proto.FavoriteStatus{UserId: user_id, VideoId: video.VideoId, IsFavorite: false})
		}

		//用协程去调用两个微服务，批量查询user实体和favoriteStatus实体
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			//调用usersinfo方法，查一批User实体
			users, _ = rpc_server.GetUsersInfo(userIds, req.Token)
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
				if user_id == -1 { //没登陆，不用查了
					break
				}
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
	}

	//返回
	resp.StatusCode = 0
	resp.StatusMsg = "查询视频成功"
	resp.VideoList = videoResult

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
		Id:             user.Id,
		Name:           user.Name,
		FollowCount:    user.FollowCount,
		FollowerCount:  user.FollowerCount,
		IsFollow:       user.IsFollow,
		TotalFavorited: user.TotalFavorited,
		WorkCount:      user.WorkCount,
		FavoriteCount:  user.FavoriteCount,
	}
}

//
//func BuildProtoVideo(item *model.Video, token string) *service.Video {
//	isFavorite := false
//	var userId int64
//	var err error
//	if token != "" { //未登录用户
//		userId, err = rpc_server.GetIdByToken(token)
//		//没有错误，说明token存在且有效，userId是解析出的当前用户id
//		if err == nil {
//			//登录状态异常，也可以刷视频，用户信息查不出来
//			isFavorite, err = rpc_server.GetFavoriteStatus(item.VideoId, userId)
//			if err != nil {
//				fmt.Println("调用远程favorite服务失败,错误原因是：")
//				fmt.Println(err)
//				return &service.Video{}
//			}
//		}
//	} else {
//		userId = 0
//	}
//
//	video := service.Video{
//		Id:            item.VideoId,
//		Author:        BuildProtoUser(item.UserId, token),
//		PlayUrl:       item.PlayUrl,
//		CoverUrl:      item.CoverUrl,
//		FavoriteCount: item.FavoriteCount,
//		CommentCount:  item.CommentCount,
//		IsFavorite:    isFavorite,
//		Title:         item.Title,
//	}
//	return &video
//}
//
//func BuildProtoUser(item_id int64, token string) *service.User {
//	rpcUserInfo, err := rpc_server.GetUserInfo(item_id, token)
//	if err != nil {
//		fmt.Println("调用远程user服务失败,错误原因是：")
//		fmt.Println(err)
//		return &service.User{}
//	}
//	//如果是空，没登陆，返回的应该是默认值
//	if rpcUserInfo == nil {
//		return &service.User{}
//	}
//	user := service.User{
//		Id:            rpcUserInfo.Id,
//		Name:          rpcUserInfo.Name,
//		FollowCount:   rpcUserInfo.FollowCount,
//		FollowerCount: rpcUserInfo.FollowerCount,
//		IsFollow:      rpcUserInfo.IsFollow,
//	}
//	return &user
//}
