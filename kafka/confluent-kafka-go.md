##error
$ go build t.go 
# pkg-config –cflags rdkafka 
Package rdkafka was not found in the pkg-config search path. 
Perhaps you should add the directory containing `rdkafka.pc’ 
to the PKG_CONFIG_PATH environment variable 
No package ‘rdkafka’ found 
pkg-config: exit status 1
--------------------- 
解决办法：
文件~/.bashrc 末尾添加

export PKG_CONFIG_PATH=/usr/lib/pkgconfig

使之生效：

$ source ~/.bashrc

