package main

import "fmt"

type ServerNode struct {
	Address      string
	Weight       int //  服务设计的权重
	CruteWeight  int //  现在的权重
	EffectWeight int //  受影响的权重
}

// 服务器设计
type Servers struct {
	Server []*ServerNode
}

func NewServers() *Servers {
	return &Servers{}
}

// 追加服务
func (server *Servers) Append(node ServerNode) {
	server.Server = append(server.Server, &node)
}

// 初始化相关的服务
func initServer() *Servers {
	server := NewServers()
	server.Append(ServerNode{Address: "www.baidu.com", Weight: 3, CruteWeight: 3, EffectWeight: 3})
	server.Append(ServerNode{Address: "www.albaba.com", Weight: 2, CruteWeight: 2, EffectWeight: 2})
	server.Append(ServerNode{Address: "www.tencent.com", Weight: 1, CruteWeight: 1, EffectWeight: 1})
	return server
}

// reblance实现负载均衡
func (server *Servers) Reblance() *ServerNode {
	// 节点总权重
	total := 0
	// 匹配获取到的节点
	var maxNode *ServerNode
	for i := 0; i < len(server.Server); i++ {
		node := server.Server[i]
		total += node.Weight
		node.CruteWeight = node.CruteWeight + node.EffectWeight
		if maxNode == nil || node.CruteWeight > maxNode.CruteWeight {
			maxNode = node
		}
		// 出现错误的时候 EffectWeight -1 正常后加 +1 权重不能超过weight
	}

	maxNode.CruteWeight = maxNode.CruteWeight - total
	return maxNode
}

// 1 设置为当前的权重 =》当前的权重+ 受影响的权重
// 2 获取当前权重最高的对象
// 3 权重最高的对象 =》 当前的权重 = 当前的权重 - 所有服务的权重之和

func main() {
	server := initServer()
	for i := 0; i < 30; i++ {
		node := server.Reblance()
		fmt.Println(node)
	}
}
