package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	fmt.Println("您好.请输入kafka地址")
	var brokers string = ""
	fmt.Scanln(&brokers)
	if brokers == "" {
		brokers = "192.168.1.246:9092"
	}

	// topic := "acq_sdk_data_1"
	myka := NewMyKafkaMsg(brokers)
	for {
	TOPIC:
		fmt.Println("请输入消费topic")
		topic := "acq_sdk_data_1"
		fmt.Scanln(&topic)
		if topic == "" {
			goto TOPIC
		}
		myka.Topic = topic
		fmt.Println("请输入group")

		group := ""
		fmt.Scanln(&group)
		myka.Group = group
	ACTION:
		action := ""
		fmt.Println("请输入动作 1 列出topic 2 列出offset 0 退出 3 重置 topic 4 重置 group ")
		fmt.Scanln(&action)
		switch action {
		case "1":
			myka.ListTopics()
		case "2":
			myka.ListOffset()
		case "3":
			t := ""
			fmt.Println("请输入新的topic")
			fmt.Scanln(&t)
			if t != "" {
				myka.Topic = t
				fmt.Println("topic重置成功")
			} else {
				fmt.Println("修改topic失败")
			}
		case "4":
			t := ""
			fmt.Println("请输入新的group")
			fmt.Scanln(&t)
			if t != "" {
				myka.Group = t
				fmt.Println("group重置成功")
			} else {
				fmt.Println("修改group失败")
			}
		case "0":
			return
		}
		goto ACTION
	}

	defer myka.Close()
}
func NewMyKafkaMsg(broker string) *MyKafkamsg {
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_2_1
	client, err := sarama.NewClient(strings.Split(broker, ","), config)
	if err != nil {
		fmt.Printf("metadata_test try create client err :%s\n\n", err.Error())
		return nil
	}
	return &MyKafkamsg{
		Client: client,
	}
}

type MyKafkamsg struct {
	Client sarama.Client
	Topic  string
	Group  string
}

func (cli *MyKafkamsg) Close() {
	cli.Client.Close()
}

func (cli *MyKafkamsg) GetPations() []int32 {
	pations, err := cli.Client.Partitions(cli.Topic)
	if err != nil {
		return nil
	}
	return pations
}

func (cli *MyKafkamsg) ListTopics() {
	client := cli.Client
	// get topic set
	topics, err := client.Topics()
	if err != nil {
		fmt.Printf("try get topics err %s\n\n", err.Error())
		return
	}

	fmt.Printf("topics(%d):\n", len(topics))
	for _, topic := range topics {
		fmt.Println(topic)
	}
}

func (cli *MyKafkamsg) GetTopicOffset() {
	client := cli.Client
	topic := cli.Topic
	pations, err := client.Partitions(topic)
	if err != nil {
		return
	}
	for _, v := range pations {
		offset, err := client.GetOffset(topic, v, time.Now().Unix())
		if err != nil {
			fmt.Println("get offset err=>", err.Error())
		}
		fmt.Printf("topic:%s pation:%d offset:%d \n\n", topic, v, offset)
	}
}

func (cli *MyKafkamsg) ListOffset() {
	client := cli.Client
	topic := cli.Topic
	group := cli.Group
	if group == "" {
		for index := range client.Brokers() {
			bk, bkerr := client.Broker(int32(index))
			if bkerr != nil {
				fmt.Println("listoffset", bkerr.Error())
				continue
			}
			gp := sarama.ListGroupsRequest{}
			res, err := bk.ListGroups(&gp)
			if err != nil {
				fmt.Println(err.Error())
			}
			for k, _ := range res.Groups {
				pations, _ := client.Partitions(topic)
				for _, v := range pations {
					offset := GetOffset(topic, k, v, client)
					fmt.Printf("topic: %s group:%s pation:%d offset:%d \n\n", topic, k, v, offset)
				}
			}
		}
	} else {
		pations, _ := client.Partitions(topic)
		fmt.Println("分区", pations)
		for _, v := range pations {
			offset := GetOffset(topic, group, v, client)
			fmt.Printf("topic: %s group:%s pation:%d offset:%d \n\n", topic, group, v, offset)
		}
	}
}

// list group
func (cli *MyKafkamsg) ListGroups() {
	client := cli.Client
	for index := range client.Brokers() {
		bkk, _ := client.Broker(int32(index))
		gp := sarama.ListGroupsRequest{}
		res, err := bkk.ListGroups(&gp)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for k, _ := range res.Groups {
			fmt.Println("group:", k)
		}
	}
}

// 获取当前的偏移量
func GetOffset(topic, group string, partion int32, client sarama.Client) int64 {
	m, _ := sarama.NewOffsetManagerFromClient(group, client)
	mm, err := m.ManagePartition(topic, partion)
	f, _ := mm.NextOffset()
	if err != nil {
		return 0
	}
	return f
}
