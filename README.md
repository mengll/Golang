# Golang
###切片清空  	chans = chans[:0]
###interface 类型断言 interface{}(a).(string) interface{}类型转化下，在进行断言

##对于一个数据类型的断言
var jk = "12"

	switch interface{}(jk).(type) {
	case string:
		fmt.Println("This is 啊string ")
	}

