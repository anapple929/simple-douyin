package main

import (
	"favorite/core"
	"favorite/core/tovideo"
	etcdInit "favorite/etcd"
	proto "favorite/service"
	utils "favorite/utils/redis"
	"github.com/micro/go-micro/v2"
)

func main() {
	//redis
	utils.InitRedis()
	etcdReg := etcdInit.EtcdReg
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcFavoriteService"), // 微服务名字
		micro.Address("127.0.0.1:8086"),
		micro.Registry(etcdReg), // etcd注册件
		micro.Metadata(map[string]string{"protocol": "http"}),
	)

	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = proto.RegisterToVideoServiceHandler(microService.Server(), new(tovideo.ToVideoService))
	_ = proto.RegisterFavoriteServiceHandler(microService.Server(), new(core.FavoriteService))
	// 启动微服务

	_ = microService.Run()

}
