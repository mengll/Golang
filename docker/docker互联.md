```
docker 的幻想链接
 docker run --link=[CONTAINER]:[ALIAS] [IMAGE] [COMMAND]

docker run -it --name test1 --link=test2:web centos
// 根据centos镜像启动一个名为test1的容器并将连接到test2上的连接起一个别名web。这样在启动test1后，比如使用```ping web```就表示ping test2容器

---------------------
docker守护进程中的默认选项：--icc默认是true，表示允许所有连接，改为false即为拒绝所有连接。

允许特定容器间连接
需要三个配置：

--icc=false
--iptables=true
--link
```
