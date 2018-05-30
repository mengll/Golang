package v2
//nsq 生产者，消费者 自行处理生产，消费内容
import (
	"github.com/nsqio/go-nsq"
	"smallgamepk.qcwanwan.com/utils"
)

type Nsqer interface {
	Producer(addr string)(*nsq.Producer,error)
	Customer(addr,topic,channel string,dat chan interface{})(error)
}

type Nsq struct {}


func NewNsq() Nsqer{
	return &Nsq{}
}

/*
	if err := producer.Publish("test", []byte("test message")); err != nil {
			log.Fatal("publish error: " + err.Error())
     }
 */
func (self *Nsq) Producer(addr string)(*nsq.Producer,error){
	cfg := nsq.NewConfig()
	NsqProducer, err := nsq.NewProducer(addr, cfg)
	if err != nil {
		utils.Log(err)
		return nil,err
	}
	return NsqProducer,nil
}

func (self *Nsq) Customer(addr,topic,channel string ,data chan interface{}) error{

	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		utils.Log(err)
	}

	// 设置消息处理函数
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		data <- message
		utils.Log(string(message.Body))
		return nil
	}))

	// 连接到单例nsqd
	if err := consumer.ConnectToNSQD(addr); err != nil {
		utils.Log(err)
	}

	<-consumer.StopChan

	return nil
}
