package core

import (
	"context"
	"favorite/mapper"
	proto "favorite/service"
	"favorite/utils"
)

type FavoriteService struct {
}

var favmapper *mapper.FavoriteMapper

func init() {
	favmapper = mapper.FavoriteMapperInstance()
}
func (*FavoriteService) FavoriteAction(ctx context.Context, in *proto.DouyinFavoriteActionRequest, out *proto.DouyinFavoriteActionResponse) error {
	return nil

}

func (*FavoriteService) FavoriteList(ctx context.Context, in *proto.DouyinFavoriteListRequest, out *proto.DouyinFavoriteListResponse) error {
	_, err := utils.GetIdByToken(in.Token)
	if err != nil {
		out.StatusCode = 500
		out.StatusMsg = "登录失效"

	}
	id := in.UserId
	var videoIds []int64
	videoIds = favmapper.GetVideoIds(id)
	println(videoIds)

	return nil
}
