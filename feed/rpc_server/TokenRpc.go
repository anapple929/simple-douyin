package rpc_server

import (
	"context"
	"feed/rpc_server/etcd"
	tokenproto "feed/services/tokenproto"
	"fmt"
	"github.com/micro/go-micro/v2"
)

func GetIdByToken(token string) (int64, error) {
	//etcdReg := etcd.NewRegistry(
	//	registry.Addrs("127.0.0.1:2379"),
	//)
	//// 服务调用实例

	//tokenMicroService := micro.NewService(micro.Name("tokenService.client"), micro.Registry(etcdReg))
	tokenMicroService := micro.NewService(micro.Registry(etcdInit.EtcdReg))
	tokenService := tokenproto.NewTokenService("rpcTokenService", tokenMicroService.Client()) //client.DefaultClient

	var req tokenproto.GetIdByTokenRequest

	req.UserToken = token

	resp, err := tokenService.GetIdByToken(context.TODO(), &req)
	if err != nil {
		fmt.Println(err)
	}
	return int64(resp.UserId), err
}
