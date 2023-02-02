package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	services "test/services"
)

func main() {

	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	//// 服务调用实例

	tokenMicroService := micro.NewService(micro.Name("tokenService.client"), micro.Registry(etcdReg))

	tokenService := services.NewTokenService("rpcTokenService", tokenMicroService.Client()) //client.DefaultClient

	var req services.GetIdByTokenRequest
	fmt.Println("到这里了")

	req.UserToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NDksImV4cCI6MTY3NTI3NTU3OSwiaXNzIjoiMTEyMjIzMyJ9.oiBkALbk8qJnIbUOZXzO9oEulKqwxeabVQ1b2VCEOJM"

	resp, err := tokenService.GetIdByToken(context.TODO(), &req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.UserId)

}
