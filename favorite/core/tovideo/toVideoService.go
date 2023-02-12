package tovideo

import (
	"context"
	"favorite/mapper"
	proto "favorite/service"
	redis "favorite/utils/redis"
	"fmt"
	"log"
	"strconv"
)

type ToVideoService struct {
}

func (*ToVideoService) GetFavoriteStatus(ctx context.Context, in *proto.GetFavoriteStatus_Request, out *proto.GetFavoriteStatus_Response) error {
	vid := in.VideoId
	uid := in.UserId
	status, err := mapper.FavoriteMapper{}.GetFavoriteStatus(vid, uid)
	if err != nil {
		return err
	}
	out.IsFavorite = status
	return nil
}

/**
根据传进来的{userId,videoId,isFavorite}集合，查询数据库的isFavorite后，返回{userId,videoId,isFavorite}集合
*/
func (*ToVideoService) GetFavoritesStatus(ctx context.Context, in *proto.GetFavoritesStatus_Request, out *proto.GetFavoritesStatus_Response) error {
	var result []*proto.FavoriteStatus
	for _, favoriteStatus := range in.FavoriteStatus {
		//查询一下redis有这个点赞关系记录吗？
		//构造key userid+videoid
		key := strconv.FormatInt(favoriteStatus.UserId, 10) + "+" + strconv.FormatInt(favoriteStatus.VideoId, 10)
		count, err := redis.RdbUserVideo.Exists(redis.Ctx, key).Result()
		if err != nil {
			log.Println(err)
		}
		if count > 0 { //缓存里有
			//redis，从redis中查是否有点赞关系
			isFavoriteRedis, err := redis.RdbUserVideo.Get(redis.Ctx, key).Result()
			if err != nil { //若查询缓存出错，则打印log
				//return 0, err
				log.Println("调用redis查询userId对应的信息出错", err)
			}
			status, _ := strconv.ParseBool(isFavoriteRedis)
			result = append(result, &proto.FavoriteStatus{UserId: favoriteStatus.UserId, VideoId: favoriteStatus.VideoId, IsFavorite: status})
		} else {
			fmt.Println("查数据库")
			status, _ := mapper.FavoriteMapper{}.GetFavoriteStatus(favoriteStatus.VideoId, favoriteStatus.UserId)
			res := &proto.FavoriteStatus{UserId: favoriteStatus.UserId, VideoId: favoriteStatus.VideoId, IsFavorite: status}
			result = append(result, res)
			//把查到的数据放入redis
			_ = redis.RdbUserVideo.Set(redis.Ctx, key, strconv.FormatBool(status), 0).Err()
		}
	}
	fmt.Println(result)
	out.IsFavorite = result
	out.StatusCode = 0
	return nil
}
