package rpc_server

import (
	"context"
	"feed/rpc_server/etcd"
	services "feed/services/favorite_to_video_proto"
	"fmt"
	"github.com/micro/go-micro/v2"
)

func GetFavoriteStatus(videoId int64, userId int64) (bool, error) {
	//// 服务调用实例
	MicroService := micro.NewService(micro.Registry(etcdInit.EtcdReg))
	Service := services.NewToVideoService("rpcFavoriteService", MicroService.Client())

	var req services.GetFavoriteStatus_Request
	req.VideoId = videoId
	req.UserId = userId

	resp, err := Service.GetFavoriteStatus(context.TODO(), &req)

	if err != nil {
		fmt.Println(err)
	}
	return resp.IsFavorite, err
}

/**
调用favorite服务，批量查询是否点赞接口
*/
func GetFavoritesStatus(isFavorites []*services.FavoriteStatus) ([]*services.FavoriteStatus, error) {
	//// 服务调用实例
	MicroService := micro.NewService(micro.Registry(etcdInit.EtcdReg))
	Service := services.NewToVideoService("rpcFavoriteService", MicroService.Client())

	var req services.GetFavoritesStatus_Request

	req.FavoriteStatus = isFavorites
	//req.Token = token

	resp, err := Service.GetFavoritesStatus(context.TODO(), &req)
	if err != nil {
		fmt.Println(err)
	}
	return resp.IsFavorite, err
}
