package core

import (
	"context"
	proto "favorite/service"
)

type FavoriteService struct {
}

func (*FavoriteService) FavoriteAction(ctx context.Context, in *proto.DouyinFavoriteActionRequest, out *proto.DouyinFavoriteActionResponse) error {
	return nil

}
func (*FavoriteService) FavoriteList(ctx context.Context, in *proto.DouyinFavoriteListRequest, out *proto.DouyinFavoriteListResponse) error {
	return nil
}
