数据总线 8个字节 64位
每个数据公用一个地址

字节必须是8 的倍数对其

内存对其数据边界
如果不对齐则一个数据需要多次取值

机器字长 =》 最大对其边界

获取类型大小与平台最大对齐边界中。较小的

type T struct
{           对齐值       最大对其边界
a int8      1 byte         8
b int64     8 byte
c int32     4 byte
d int16     2 byte
}

addr % 8 =0 

addr 是对齐边界的倍数

0 1  2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 2
a
0%1=0 

1%8 != 0         b                     
                                       16%4 =0  
                                                   20%2=0 
