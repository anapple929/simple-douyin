package main

import (
	"comment/core"
	proto "comment/service"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {

	// etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcCommentService"), // 微服务名字
		micro.Address("127.0.0.1:8087"),
		micro.Registry(etcdReg), // etcd注册件
		micro.Metadata(map[string]string{"protocol": "http"}),
	)

	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = proto.RegisterCommentServiceHandler(microService.Server(), new(core.CommentService))
	// 启动微服务

	_ = microService.Run()

}
