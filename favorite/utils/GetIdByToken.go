package utils

import (
	"context"
	"favorite/etcd"
	proto "favorite/service"
	"fmt"
	"github.com/micro/go-micro/v2"
)

func GetIdByToken(token string) (int64, error) {
	if "" == token {
		return -1, nil
	}

	tokenMicroService := micro.NewService(micro.Registry(etcdInit.EtcdReg))

	tokenService := proto.NewTokenService("rpcTokenService", tokenMicroService.Client()) //client.DefaultClient

	var req proto.GetIdByTokenRequest

	req.UserToken = token

	resp, err := tokenService.GetIdByToken(context.TODO(), &req)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	return int64(resp.UserId), nil
}
