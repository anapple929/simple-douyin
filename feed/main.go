package main

import (
	"context"
	services "feed/services"
<<<<<<< HEAD
=======
	"feed/wrappers"
>>>>>>> 7d37cac7c00e284ee5031d5d3ad2cad9d3258d4c
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
<<<<<<< HEAD
	tokenMicroService := micro.NewService(micro.Name("tokenService.client"), micro.Registry(etcdReg))
=======
	tokenMicroService := micro.NewService(micro.Name("tokenService.client"),
		micro.WrapClient(wrappers.NewTokenWrapper), micro.Registry(etcdReg))
>>>>>>> 7d37cac7c00e284ee5031d5d3ad2cad9d3258d4c

	tokenService := services.NewTokenService("rpcTokenService", tokenMicroService.Client()) //client.DefaultClient

	var req services.GetIdByTokenRequest
	fmt.Println("到这里了")

<<<<<<< HEAD
	req.UserToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NDksImV4cCI6MTY3NTI3NTU3OSwiaXNzIjoiMTEyMjIzMyJ9.oiBkALbk8qJnIbUOZXzO9oEulKqwxeabVQ1b2VCEOJM"
=======
	req.UserToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NDMsImV4cCI6MTY3NTI2NzEzMywiaXNzIjoiMTEyMjIzMyJ9.1PVRVAMbax8-gTT0L8pl-72MmMU5dAITs7H-ylY2uVs"
>>>>>>> 7d37cac7c00e284ee5031d5d3ad2cad9d3258d4c

	resp, err := tokenService.GetIdByToken(context.TODO(), &req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.UserId)

}
