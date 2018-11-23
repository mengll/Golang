package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"google.golang.org/grpc/naming"
)


func main() {

	// 创建服务
	cli, err := clientv3.New(clientv3.Config{
		// etcd集群成员节点列表
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	// 创建命名解析 服务注册的实现 注册grpcf 注册命名解析
	dns := &etcdnaming.GRPCResolver{Client: cli}
	dns.Update(context.TODO(), "myserver", naming.Update{Op: naming.Add, Addr: "127.0.0.1:3000"})

}
