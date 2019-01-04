package fuck

import "fmt"

var Mp map[string]Events

//各种媒体渠道注册
func init() {
	Mp = make(map[string]Events)
	Mp["func"] = &Event{}
	Mp["func"].Construct()
	Mp["func"].Report("today", map[string]string{"fuck": "you"})
}

// 定义事件处理的过程
type Event struct{}

type Events interface {
	Report(name string, data map[string]string) error
	Construct()
}

// 执行上报的操作
func (self *Event) Report(name string, data map[string]string) error {
	fmt.Println(name, data)
	return nil
}

// 数据初始化的操作
func (self *Event) Construct() {
	fmt.Println("init data")
}

func Fuck(b *bool) {
	select {
	default:
		*b = true
	}
}
