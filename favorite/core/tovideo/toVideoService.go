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
