package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	//services "test/services/to_favorite"
	//video "test/services/to_video"
	fav "test/services/fav"
)

func main() {
	//测试第一个功能

	//etcdReg := etcd.NewRegistry(
	//	registry.Addrs("127.0.0.1:2379"),
	//)
	//
	////// 服务调用实例
	//MicroService := micro.NewService(micro.Registry(etcdReg))
	//Service := services.NewToFavoriteService("rpcPublishService", MicroService.Client()) //client.DefaultClient
	//
	//var req services.UpdateFavoriteCountRequest
	//fmt.Println("到这里了")
	//
	//req.VideoId = 30
	//req.Count = 1
	//req.Type = 1
	//
	//resp, err := Service.UpdateFavoriteCount(context.TODO(), &req)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(resp.StatusCode)
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	//测试第二个功能
	//// 服务调用实例
	/*	MicroService := micro.NewService(micro.Registry(etcdReg))
		Service := services.NewToFavoriteService("rpcPublishService", MicroService.Client()) //client.DefaultClient*/

	/*var req services.GetVideosByIdsRequest

	var videoId []int64
	videoId = append(videoId, 26)
	videoId = append(videoId, 27)
	videoId = append(videoId, 28)
	videoId = append(videoId, 30)
	req.VideoId = videoId

	resp, err := Service.GetVideosByIds(context.TODO(), &req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.VideoList)*/
	MicroService := micro.NewService(micro.Registry(etcdReg))
	Service := fav.NewFavoriteService("rpcFavoriteService", MicroService.Client())
	var req fav.DouyinFavoriteActionRequest
	req.VideoId = 39
	req.ActionType = 1
	req.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NTEsImV4cCI6MTY3NTQxNzk5MiwiaXNzIjoiMTEyMjIzMyJ9.xaRQZZp-ecCg4dNkXMWiNMWgF3Wv8uKr_HvQxIqmxQQ"
	status, err := Service.FavoriteAction(context.TODO(), &req) //client.DefaultClient.
	if err != nil {
		fmt.Println("err====", err)
	}
	fmt.Println(status)
}
