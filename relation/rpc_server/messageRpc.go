package rpc_server

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	etcdInit "relation/rpc_server/etcd"
	from_message "relation/service/from_message"
)

/*
*
输入QueryBody列表，查询Message实体列表
*/
func QueryMessagesByUsers(queryBody []*from_message.QueryBody) ([]*from_message.QueryBody, error) {
	MicroService := micro.NewService(micro.Registry(etcdInit.EtcdReg))
	Service := from_message.NewToRelationService("rpcMessageService", MicroService.Client())

	var req from_message.QueryMessagesByUsersRequest
	req.QueryBody = queryBody

	resp, err := Service.QueryMessagesByUsers(context.TODO(), &req)
	if err != nil {
		fmt.Println("调用远程Message服务失败，具体错误如下")
		fmt.Println(err)
	}
	fmt.Println("调用回来了")
	fmt.Println(resp.QueryBody)

	return resp.QueryBody, err
}
