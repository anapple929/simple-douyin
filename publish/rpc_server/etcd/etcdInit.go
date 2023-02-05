package etcdInit

import (
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

var EtcdReg registry.Registry

func init() {
	EtcdReg = etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
}
