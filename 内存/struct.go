1) struct 的内存分布原理
原理
关于Golang同一struct中field的书写顺序不同内存分配大小也会不同。主要原因如下：struct内field内存分配是以4B为基础，超过4B时必须独占。

type A1 struct {
    a bool
    b uint32
    c bool
    d uint32
    e uint8
    f uint32
    g uint8
}

计算一下A1所需要占用的内存：

首先第1个4B中放入a，a是bool型，占用1B，剩余3B
这时看b是uint32，占用4B，剩余3B放不下，所以offset到下一个4B空间，这时我们会发现3B没有放东西，被浪费了
依次往下，A1要占用28B的空间
根据1，2两个步骤很容易看出，有很多浪费空间。

优化：
type A2 struct {
    a bool
    c bool
    e uint8
    g uint8
    b uint32
    d uint32
    f uint32
}

首先第1个4B中放入a，a是bool型，占用1B，剩余3B
c是bool，占用1B，放入后剩余2B
d是uint8，占用1B，放入后剩余1B
依次往下
这样会使内存使用率高很多。

原文地址 https://studygolang.com/articles/13364?fr=sidebar
