package tofavorite

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"publish/model"
	"publish/rpc_server"
	"publish/services/favorite_to_video_proto"
	proto "publish/services/to_favorite"
	usersproto "publish/services/to_relation"
	redis "publish/utils"
	"strconv"
	"sync"
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

	var searchIds []int64          //保存从redis没命中的videoId,统一去mysql中查一次
	var videosRedis []*model.Video //从redis查到的video实体
	//把能从redis中查到的videoId先都给查到，查不到的放到searchIds中，在数据库中统一给查出来
	for _, videoIdSearchRedis := range req.VideoId {
		count, err := redis.RdbVideoId.Exists(redis.Ctx, strconv.FormatInt(videoIdSearchRedis, 10)).Result()
		if err != nil {
			log.Println(err)
		}

		if count > 0 {
			var video *model.Video //一个临时变量，存解构后的video实体
			//缓存里有
			//redis，先从redis通过videoId查询video实体
			userString, err := redis.RdbVideoId.Get(redis.Ctx, strconv.FormatInt(videoIdSearchRedis, 10)).Result()
			if err != nil { //若查询缓存出错，则打印log
				//return 0, err
				log.Println("调用redis查询videoId对应的信息出错", err)
			}
			json.Unmarshal([]byte(userString), &video)
			videosRedis = append(videosRedis, video)
		} else {
			//没查到，进入searchIds数组
			searchIds = append(searchIds, videoIdSearchRedis)
		}
	}

	//调用数据库查video实体列表
	videos, err := model.NewVideoDaoInstance().GetVideosByIds(searchIds)
	//把mysql查到的videos数据存到redis中
	for _, video := range videos {
		videoValue, _ := json.Marshal(&video)
		_ = redis.RdbVideoId.Set(redis.Ctx, strconv.FormatInt(video.VideoId, 10), videoValue, 0).Err()
	}

	if err != nil {
		resp.StatusCode = -1
		resp.VideoList = nil
		return errors.New("调用数据库出错！")
	}

	//合并从mysql查到的users和redis查到的usersRedis
	videos = append(videos, videosRedis...)

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
	wg.Add(2) //等待两个协程都拿到数据了，再往下走

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
