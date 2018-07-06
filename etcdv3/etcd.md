```
集群配置 需要关闭防火墙 不然无法里链接，或者开启防火墙
```

  
etcd --name infra0 --initial-advertise-peer-urls http://192.168.5.130:2380 \
--listen-peer-urls http://192.168.5.130:2380 \
--listen-client-urls http://192.168.5.130:2379,http://127.0.0.1:2379 \
--advertise-client-urls http://192.168.5.130:2379 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster infra0=http://192.168.5.130:2380,infra1=http://192.168.5.131:2380,infra2=http://192.168.5.132:2380 \
--initial-cluster-state new


etcd --name infra1 --initial-advertise-peer-urls http://192.168.5.131:2380 \
--listen-peer-urls http://192.168.5.131:2380 \
--listen-client-urls http://192.168.5.131:2379,http://127.0.0.1:2379 \
--advertise-client-urls http://192.168.5.131:2379 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster infra0=http://192.168.5.130:2380,infra1=http://192.168.5.131:2380,infra2=http://192.168.5.132:2380 \
--initial-cluster-state new
  
etcd --name infra2 --initial-advertise-peer-urls http://192.168.5.132:2380 \
--listen-peer-urls http://192.168.5.132:2380 \
--listen-client-urls http://192.168.5.132:2379,http://127.0.0.1:2379 \
--advertise-client-urls http://192.168.5.132:2379 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster infra0=http://192.168.5.130:2380,infra1=http://192.168.5.131:2380,infra2=http://192.168.5.132:2380 \
--initial-cluster-state new
  
