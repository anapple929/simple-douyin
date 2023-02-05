package etcdInit

import (
	"context"
	proto "favorite/service"
	"favorite/service/frompublish"

	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

var EtcdReg registry.Registry

func init() {
	EtcdReg = etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

}
func CountAction(vid int64, count int32, actionType int32) bool {
	MicroService := micro.NewService(micro.Registry(EtcdReg))
	Service := frompublish.NewToFavoriteService("rpcPublishService", MicroService.Client())
	var req frompublish.UpdateFavoriteCountRequest
	req.VideoId = vid
	req.Count = count
	req.Type = actionType
	favoriteCount, err := Service.UpdateFavoriteCount(context.TODO(), &req)
	if err != nil || favoriteCount.StatusCode != 0 {
		fmt.Println("favoriteCount维护失败:", err)
		return false
	}
	return true

}

func GetVideosByIds(vids []int64) ([]*proto.Video, error) {
	//// 服务调用实例

	MicroService := micro.NewService(micro.Registry(EtcdReg))
	Service := frompublish.NewToFavoriteService("rpcPublishService", MicroService.Client()) //client.DefaultClient

	var req frompublish.GetVideosByIdsRequest

	req.VideoId = vids
	resp, err := Service.GetVideosByIds(context.TODO(), &req)
	if err != nil {
		fmt.Println("远程调用错误", err)
		return nil, err
	}

	return changeVideo(resp.VideoList), nil

}
func changeVideo(videos []*frompublish.Video) []*proto.Video {
	var res []*proto.Video

	for i := 0; i < len(videos); i++ {
		video := videos[i]
		one := proto.Video{
			Id:            video.Id,
			Author:        changeUser(video.Author),
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		}
		res = append(res, &one)
	}
	return res
}

func changeUser(fuser *frompublish.User) *proto.User {
	return &proto.User{
		Id:            fuser.Id,
		Name:          fuser.Name,
		FollowCount:   fuser.FollowCount,
		FollowerCount: fuser.FollowerCount,
		IsFollow:      fuser.IsFollow,
	}

}
