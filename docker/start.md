### docker 执行挂在 端口
```
docker.exe run -it -p 8080:1323 -v e:/Go/tool/src:/go/src mllgo:v1 /go/src/pk_admin/main

-p 宿主机ip：容器ip
-v 卷 宿主机 ：容器地址
镜像名：版本
要执行的命令
 顺序可能影响执行结果
 
docker.exe build -f .\Dockerfile -t mll\go . 在当前目录下   

历史容器查看 
docker ps -n 10 展示最近的是个容器 
docker ps -a 展示所用 -q 展示容器号

# 新创建一个容器并运行
docker run 

#运行一个已经存在的容器
docker start 

# docker attach 重新的回归导已运行的容器上 还有 docker exec -it 

```
