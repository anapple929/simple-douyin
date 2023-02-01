package main

import (
	"context"
	services "feed/services"
	"feed/wrappers"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {

	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	//// 服务调用实例
	tokenMicroService := micro.NewService(micro.Name("tokenService.client"),
		micro.WrapClient(wrappers.NewTokenWrapper), micro.Registry(etcdReg))

	tokenService := services.NewTokenService("rpcTokenService", tokenMicroService.Client()) //client.DefaultClient

	var req services.GetIdByTokenRequest
	fmt.Println("到这里了")

	req.UserToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NDMsImV4cCI6MTY3NTI2NzEzMywiaXNzIjoiMTEyMjIzMyJ9.1PVRVAMbax8-gTT0L8pl-72MmMU5dAITs7H-ylY2uVs"

	resp, err := tokenService.GetIdByToken(context.TODO(), &req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.UserId)

}
