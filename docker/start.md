### docker 执行挂在 端口
```
docker.exe run -it -p 8080:1323 -v e:/Go/tool/src:/go/src mllgo:v1 /go/src/pk_admin/main

-p 宿主机ip：容器ip
-v 卷 宿主机 ：容器地址
镜像名：版本
要执行的命令
 顺序可能影响执行结果
```
