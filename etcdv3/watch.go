```
读取键值的变化
```

package watch

import (
	"github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
	"context"
)

func init(){
	go watch()
}

func watch(){
	cli,err :=clientv3.New(clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
		DialTimeout: time.Duration(5) * time.Second, //链接超时时间
	})

	if err != nil{
		fmt.Println(err.Error())
		return
	}
	//ctx_a,cancel_a := context.WithTimeout(context.Background(),time.Second*10) //设置读取超时时间
	//
	//p,e := cli.Grant(context.Background(),int64(13))
	//
	//if e != nil{
	//	fmt.Println(e.Error())
	//}
	//
	//fmt.Println(p.ID)
	// p.ID 可视化编辑工具

	for {
		select{
		case item := <-cli.Watch(context.Background(),"/mll",clientv3.WithPrefix()):
			for _,v := range item.Events {
				fmt.Println("数据改变了打的撒旦",string(v.Kv.Value))
			}

		}
	}
	//cancel_a() 取消配置的管理的方法 监测不到数据变化 
}
