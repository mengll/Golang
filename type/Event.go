```
package v2

import (
	"sync"
	"smallgamepk.qcwanwan.com/utils"
	"time"
)

type Enenter interface {
	AddEventListerner(event_name string ,back_func EventCallBack)
}

//事件处理模式
type Event struct {
	EventName  string
	Params     map[string]interface{}

}

//注册回调函数
type EventCallBack func(event *Event)

type MyEvent struct {
	EventList map[string]EventCallBack  // 注册时间和回调函数
	Events chan *Event
	M sync.RWMutex
}

//抛出事件
func(self *MyEvent) DisPatcherEvent(event *Event) error{
	defer func() {
		if err := recover();err != nil{
			utils.Log(err)
		}
	}()
	self.Events <- event
	return nil
}

//注册事件
func(self *MyEvent) AddEventListener(event_name string,back_func EventCallBack) error{
	//判断是否已经存在
	self.M.Lock()
	if _,ok := self.EventList[event_name];!ok{
		self.EventList[event_name] = back_func
	}
	self.M.Unlock()
	return nil
}

//注销事件
func (self *MyEvent) RemoveEventListener(event_name string,back_func EventCallBack) error {
	self.M.Lock()
	if _,ok := self.EventList[event_name];ok{
		delete(self.EventList,event_name)
	}
	self.M.Unlock()
	//todo 暂停方法的执行并
	return nil
}

//创建事件对象
func (self *MyEvent) NewEvent(event_name string,param map[string]interface{})(*Event,error){
	return &Event{EventName:event_name,Params:param} ,nil
}

//底层事件处理
func (self *MyEvent) EventLoop(){

	for item := range self.Events {
		event_name := item.EventName
		ck := self.EventList[event_name]
		ck(item)
	}

}

//创建数
func NewMyEvent() *MyEvent{
	return &MyEvent{
		EventList: make(map[string]EventCallBack),
		Events:make(chan *Event,100), //创建带有缓冲区的对象
	}
}

//使用实力
func Fk(event *Event) {
	utils.Log(event)
}

//事件驱动调用
func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	m := NewMyEvent()
	go m.EventLoop()
	m.AddEventListener("fk", Fk)
	mp := make(map[string]interface{})
	mp["age"] = 23
	event, _ := m.NewEvent("fk", mp)

	m.DisPatcherEvent(event)

	select {
		case <-time.After(time.Second * 3):
			wg.Done()
	}
	wg.Wait()
}






```
