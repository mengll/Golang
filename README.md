# Golang
###切片清空  	chans = chans[:0]
###interface 类型断言 interface{}(a).(string) interface{}类型转化下，在进行断言

##对于一个数据类型的断言
var jk = "12"

	switch interface{}(jk).(type) {
	case string:
		fmt.Println("This is 啊string ")
	}

``` 

type ty struct {
   Name string `json:"name"`
}
`` 部分被称为标记tag  表示的是json 解码的元数据 用于创建ty数据类型原值的切片 使用这个函数读取数据
```
