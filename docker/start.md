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

# docker top 容器  查看当前容器运行的所有的进程

# docker 调用数据卷容器
docker run -it --name mygo --volumes-from=web_c -e GOPATH=/data/golang -e GOBIN=/data/golang/bin golang /bin/bash

```

--------- win1o docker toolbox 下的操作-------
docker-machine.exe  ssh default  进入命令修改的地方

sudo /etc/init.d/docker restart 重启

sudo vi /var/lib/boot2docker/profile 配置文件

-------- 修改数据源-----------------

---------win10挂载宿主机的问题
调整盘符 完美挂在
d:\work\docker\datasource
首先将当前的路径挂载到 虚拟机的共享文件夹中  virtualbox 中的共享文件夹中,在其中的名声为 datasource

docker run -it -v /datasource:/Engine -p 8000:8000 steveny/predictionio:0.12.0 /bin/bash




