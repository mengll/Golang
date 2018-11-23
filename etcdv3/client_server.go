package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"test/chat"
	"time"

	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"google.golang.org/grpc"
)

const TS int = 0

func main() {
	cli, err := clientv3.New(clientv3.Config{
		// etcd集群成员节点列表
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("[测] connect etcd err:", err)
		return
	}

	r := &etcdnaming.GRPCResolver{Client: cli}
	b := grpc.RoundRobin(r)
  
	con, err := grpc.Dial("myserver", grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
	if err != nil {
		panic(err.Error())
	}

	// grpc的方式连接 不验证证书的方式
	// con, err := grpc.Dial(":3000", grpc.WithInsecure())
	// if err != nil {
	// 	fmt.Println("error", err.Error())
	// }
	client := chat.NewChatClient(con)
	ctx, _ := context.WithCancel(context.Background())
	fuck := context.WithValue(ctx, "mll", "fj")
	// 连接的时候
	// 数据写入的过程化的方法的操作的实现
	// bind the data source
	stream, err := client.BidStream(fuck)
	if err != nil {
		fmt.Println(err.Error())
	}

	// close the request
	// 接受当前的数据流的信息
	go func() {
		for {
			// 接受数据流的信息
			res, ers_err := stream.Recv()

			if ers_err != nil {
				fmt.Println("read message error")
				// read the data error
				break
			}
			// 做操
			if ers_err == io.EOF {
				break
			}
			fmt.Printf("%s", res)
		}

	}()

	tk := time.NewTicker(time.Duration(10) * time.Second)
	defer tk.Stop()

	for {
		select {
		case <-tk.C:
			stream.Send(&chat.Request{Input: "this is the best request"})
		case <-fuck.Done():
			log.Printf("[server close]%s", "now")
			return
		}
	}
}
