package rpc

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	tokenproto "publish/services/tokenproto"
)

func GetIdByToken(token string) int64 {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	tokenMicroService := micro.NewService(micro.Name("tokenService.client"), micro.Registry(etcdReg))

	tokenService := tokenproto.NewTokenService("rpcTokenService", tokenMicroService.Client()) //client.DefaultClient

	var req tokenproto.GetIdByTokenRequest

	req.UserToken = token

	resp, err := tokenService.GetIdByToken(context.TODO(), &req)
	if err != nil {
		fmt.Println(err)
	}
	return int64(resp.UserId)
}
