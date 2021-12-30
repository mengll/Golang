```
 1 。当翻倍的后的数据大于 最小的数据值
 
 a := []string{"onw","two","three"}
 a = append(a,"for","five")
 
 old_cap = 3
 
 want_cap = 5
 
 3*2 >5 and 3 < 1024
 
 new_cap = old_cap * 2


desc：
原始的切片长度为 3位 使用append的方式最少需要扩展到5 

原始的数据扩展后 大于最少需要的空间，然后判断当前的，数据总的长度是否大于1024 当小于1024 直接翻倍，大于1024 增加原来的 1/4

2. 翻倍小于预估计的值
a := []string{"one","two"}
a = append(a,"three","five","six")

old_cap =2 
want_cap=5

old_cap *2 < want_cap 

new_cap = want_cap
   小于之前预想的值，数据直接翻倍




```
