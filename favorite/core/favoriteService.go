package core

import (
	"context"
	etcdInit "favorite/etcd"
	"favorite/mapper"
	proto "favorite/service"
	"favorite/utils"
	"fmt"
)

type FavoriteService struct {
}

var favmapper *mapper.FavoriteMapper

func init() {
	favmapper = mapper.FavoriteMapperInstance()
}
func (*FavoriteService) FavoriteAction(ctx context.Context, in *proto.DouyinFavoriteActionRequest, out *proto.DouyinFavoriteActionResponse) error {
	actionType := in.ActionType //1点赞  2取消
	vid := in.VideoId
	token := in.Token
	uid, err := utils.GetIdByToken(token)
	if err != nil {
		out.StatusCode = 500
		out.StatusMsg = "登录失效"
		return err
	}
	err = favmapper.FavoriteAction(uid, vid, actionType, token)
	if err != nil {
		out.StatusCode = 500
		out.StatusMsg = "操作失败"
		return err
	}
	out.StatusMsg = "操作成功"
	out.StatusCode = 0
	return nil

}

func (*FavoriteService) FavoriteList(ctx context.Context, in *proto.DouyinFavoriteListRequest, out *proto.DouyinFavoriteListResponse) error {
	fmt.Println("====进入====")
	_, err := utils.GetIdByToken(in.Token)
	if err != nil {
		out.StatusCode = 500
		out.StatusMsg = "登录失效"
		return err
	}
	id := in.UserId
	var videoIds []int64
	videoIds = favmapper.GetVideoIds(id)
	fmt.Println("videoIds==", videoIds)
	resp, err := etcdInit.GetVideosByIds(videoIds, in.Token)
	if err != nil {
		fmt.Println("出错了，....")
		out.StatusCode = 500
		out.StatusMsg = "获取视频失败"
		return err
	}

	out.VideoList = resp
	out.StatusCode = 0
	out.StatusMsg = "查询成功"

	return nil
}
