package tovideo

import (
	"context"
	"favorite/mapper"
	proto "favorite/service"
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
	out.IsFavorite, _ = mapper.FavoriteMapper{}.GetFavoritesStatus(in.FavoriteStatus)
	out.StatusCode = 0
	return nil
}
