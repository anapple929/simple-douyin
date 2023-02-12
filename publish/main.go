package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"publish/conf"
	"publish/core"
	toComment "publish/core/to_comment"
	toFavorite "publish/core/tofavorite"
	"publish/services"
	protoToComment "publish/services/to_comment"
	protoToFavorite "publish/services/to_favorite"
	redis "publish/utils"
)

func main() {
	conf.Init()
	//redis
	redis.InitRedis()
	// etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcPublishService"), // 微服务名字
		micro.Address("127.0.0.1:8083"),
		micro.Registry(etcdReg), // etcd注册件
		micro.Metadata(map[string]string{"protocol": "http"}),
	)
	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = services.RegisterPublishServiceHandler(microService.Server(), new(core.PublishService))
	_ = protoToFavorite.RegisterToFavoriteServiceHandler(microService.Server(), new(toFavorite.ToFavoriteService))
	_ = protoToComment.RegisterToCommentServiceHandler(microService.Server(), new(toComment.ToCommentService))
	// 启动微服务
	_ = microService.Run()
}
