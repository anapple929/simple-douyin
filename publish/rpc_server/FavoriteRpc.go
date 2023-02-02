package rpc_server

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"publish/etcd"
	services "publish/services/favorite_to_video_proto"
)

func GetFavoriteStatus(videoId int64, userId int64) bool {

	//// 服务调用实例
	MicroService := micro.NewService(micro.Registry(etcdInit.EtcdReg))
	Service := services.NewToVideoService("rpcFavoriteService", MicroService.Client()) //client.DefaultClient

	var req services.GetFavoriteStatus_Request
	//fmt.Println("到这里了")

	req.VideoId = videoId
	req.UserId = userId

	resp, err := Service.GetFavoriteStatus(context.TODO(), &req)
	if err != nil {
		fmt.Println(err)
	}
	return resp.IsFavorite
}
