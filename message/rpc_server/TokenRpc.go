package rpc_server

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"message/rpc_server/etcd"
	tokenproto "message/service/tokenproto"
)

/**
调用token解析
*/
func GetIdByToken(token string) (int64, error) {
	tokenMicroService := micro.NewService(micro.Registry(etcdInit.EtcdReg))
	tokenService := tokenproto.NewTokenService("rpcTokenService", tokenMicroService.Client())

	var req tokenproto.GetIdByTokenRequest

	req.UserToken = token

	resp, err := tokenService.GetIdByToken(context.TODO(), &req)
	if err != nil {
		fmt.Println(err)
	}
	return int64(resp.UserId), err
}
