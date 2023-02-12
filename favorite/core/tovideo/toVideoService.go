package tovideo

import (
	"context"
	"favorite/mapper"
	proto "favorite/service"
	"fmt"
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
		status, _ := mapper.FavoriteMapper{}.GetFavoriteStatus(favoriteStatus.VideoId, favoriteStatus.UserId)
		res := &proto.FavoriteStatus{UserId: favoriteStatus.UserId, VideoId: favoriteStatus.VideoId, IsFavorite: status}
		result = append(result, res)
	}
	fmt.Println(result)
	out.IsFavorite = result
	out.StatusCode = 0
	return nil
}
