package rpc_server

import (
	etcdInit "comment/rpc_server/etcd"
	publishproto "comment/service/frompublish"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
)

func CountAction(vid int64, count int32, actionType int32) bool {
	toCommentMicroService := micro.NewService(micro.Registry(etcdInit.EtcdReg))
	toCommentService := publishproto.NewToCommentService("rpcPublishService", toCommentMicroService.Client())
	var req publishproto.UpdateCommentCountRequest
	req.VideoId = vid
	req.Count = count
	req.Type = actionType
	publishCount, err := toCommentService.UpdateCommentCount(context.TODO(), &req)
	if err != nil || publishCount.StatusCode != 0 {
		fmt.Println("commentCount维护失败:", err)
		return false
	}
	return true

}
