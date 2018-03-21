package server

import (
	"github.com/gorilla/websocket"
	"github.com/go-redis/redis"
	"fmt"
	"encoding/json"
)

const max_room_num = 2 //每个房间的最大人数

var (
	ActiveClients = make(map[string]ClientConn) //在线的用户的信息
	User          = make(map[string]string)
)

type UserMsg struct {
	Room string      `json:"room"` 			//房间名字
	Cmd string       `json:"cmd"`			//登录，初始化 准备 退出 等相关的命令
	User string	     `json:"user"`			//用户的信息
	AvatarUrl string `json:"avatar_url"` 	//用户的头像信息
	Content string	 `json:"content"` 		//用户的发送的内容
	Uuid string      `json:"uuid"`			//创建用户的唯一识别码
	Gameid string    `json:"gameid"`		//当前的游戏
}

type ClientConn struct {
	websocket *websocket.Conn
}

type UserInfo struct {
	User      string `json:"user"`  	    //玩家的用户信息
	AvatarUrl string `json:"avatar_url"`	//头像
	Uuid      string `json:"uuid"`			//用户ID
	Gameid    string `json:"gameid"`	    //游戏ID
}

type ReplyMsg struct {
	Room string		`json:"room"`           //玩家的放假
	Cmd  string		`json:"cmd"`		    //玩家执行的命令
	Data string		`json:"data"`		    //传递的数据
	Gameid string   `json:"gameid"`         //当前的游戏id
}

type DatBase struct {
	Room    string  `json:"room"`   	    // 房间号
	Command string  `json:"command"`        // 执行访问命令
	Content string  `json:"content"`        // 传递数据信息
}

type Clients struct {
	Clients UserMsg //保存用户的信息
}

type GameClient interface{
	Connect(conn *websocket.Conn)
}
var RedisClient *redis.Client

func init(){
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "192.168.1.246:6379",
		Password: "", // 设置Redis的链接的链接方法
		DB:       0,  // use default DB
	})

}

func WsConnect(conn *websocket.Conn,dat *UserMsg){

	sockCli := ClientConn{conn}
	RedisClient.Ping() //

	room := fmt.Sprintf("ROOM:%s", dat.Room) //读取房间信息

	command := dat.Cmd
	//获取房间信息
	println(room,command)

	switch command {
	case "reday":
		println("游戏退出当前登录")
		RedisClient.Set("READY:"+dat.Uuid,"ready",0)

		//从redis取房间内的所有用户uuid
		roomSlice := RedisClient.SMembers("ROOM:"+dat.Room) //更新当前的用户的状态为
		//用户uuid保存到一个go切片online
		online := roomSlice.Val()

		i := 0

		//循环取在线用户个人信息
		if len(online) != 0 {
			for _, na := range online {
				if na != "" {
					userJson := RedisClient.Get("READY:"+na)
					userJson2 := userJson.Val()
					if userJson2 == "ready" {
						i++
					}
				}
			}
		}

		if i == len(online) && i == max_room_num {
			var rm ReplyMsg
			rm.Room = dat.Room
			rm.Cmd = "start"
			rm.Data = ""
			rm.Gameid = dat.Gameid

			broadcast(RedisClient,dat,rm)
		}

	case "login":
		println("游戏进入转呗状态")

		checkNumTmp := RedisClient.SCard(room)
		checkNum := checkNumTmp.Val()
		println("房间中的人数：",checkNum)
		println(dat.Uuid)

		if(checkNum <= max_room_num) {
			fmt.Println("checkNum success")
			//socket用户列表新增当前用户websocket连接
			ActiveClients[dat.Uuid] = sockCli

			//用户uuid保存到redis房间set集合内
			RedisClient.SAdd(room, dat.Uuid)
			var userinfo UserInfo

			userinfo.User 		= dat.User
			userinfo.AvatarUrl  = dat.AvatarUrl
			userinfo.Uuid       = dat.Uuid
			userinfo.Gameid     = dat.Gameid

			//生成用户信息json串
			b, err := json.Marshal(userinfo) //格式化当前的数据信息
			if err != nil {
				fmt.Println("Encoding User Faild")
			} else {

				//保存用户信息到redis
				RedisClient.Set("USER:"+userinfo.Uuid, b, 0)
				println("保存用户信息到Redis--》")
				//初始化用户
				initOnlineMsg(RedisClient,dat)

			}
		}else {
			var rm ReplyMsg
			rm.Room = dat.Room
			rm.Cmd = "loginFailed"
			rm.Data = "登录失败，人数已满"

			sendMsg,err2 := json.Marshal(rm)
			sendMsgStr := string(sendMsg)
			fmt.Println(sendMsgStr)
			if err2 != nil {
				println("data type channge error",err2.Error())
			} else {
				if err := conn.WriteMessage(websocket.TextMessage,[]byte(sendMsgStr)); err != nil {
					println("Could not send UsersList to ", dat.User, err.Error())
				}
			}
		}

	case "logout":
		delete(ActiveClients,dat.Uuid) //删除当前用户
		println("退出当前的服务")
		RedisClient.SRem("ROOM:"+dat.Room,dat.Uuid)

		//初始化用户
		initOnlineMsg(RedisClient,dat)

	case "sendmess":
		println("发送数据")
		var rm ReplyMsg
		rm.Room = dat.Room
		rm.Cmd = "start"
		rm.Data = dat.Content
		rm.Gameid = dat.Gameid
		broadcast(RedisClient,dat,rm)  //广播当前的游戏的信息
	}

}

//房间成员初始化,有人加入或者退出都要重新初始化，相当于聊天室的在线用户列表的维护
func initOnlineMsg(redisClient *redis.Client,userMsg *UserMsg) {

	var err error

	//从redis取房间内的所有用户uuid
	roomSlice := redisClient.SMembers("ROOM:"+userMsg.Room)
	//用户uuid保存到一个go切片online
	online := roomSlice.Val()  //

	var onlineList []string

	//循环取在线用户个人信息
	if len(online) != 0 {
		for _, na := range online {
			if na != "" {
				userJson := redisClient.Get("USER:"+na)
				userJson2 := userJson.Val()
				onlineList = append(onlineList,userJson2) //获取获取房钱房间的用户的信息
			}
		}
	}

	//生成在线用户信息json在线的用户新书数据
	//c, err := json.Marshal(onlineList)

	onlineListStr,err2 := json.Marshal(onlineList)
	if err2 != nil{
		return
	}

	//数据展示的过程
	var rm ReplyMsg
	rm.Room = userMsg.Room
	rm.Cmd = "init"
	rm.Data = string(onlineListStr)

	//给所有用户发初始化消息
	if len(online) != 0 {
		for _, na := range online {
			if na != "" {
				ws := ActiveClients[userMsg.Uuid]
				if err = ws.websocket.WriteJSON(rm); err != nil {
					println("Could not send UsersList to ", "", err.Error())
				}
			}
		}
	}

	//若房间人数满，发送就绪消息
	if len(online) >= max_room_num {
		fmt.Println("full")
		var rm ReplyMsg
		rm.Room = userMsg.Room
		rm.Cmd = "full"
		rm.Data = "the game is full"

		for _, na := range online {
			if na != "" {
				ws := ActiveClients[userMsg.Uuid]
				if err = ws.websocket.WriteJSON(rm); err != nil {
					println("Could not send UsersList to ", "", err.Error())
				}
			}
		}
	}
}

//给房间内的所有用户发送相关的数据
func broadcast(redisClient *redis.Client,userMsg *UserMsg,rm ReplyMsg) {
	var err error

	//从redis取房间内的所有用户uuid
	roomSlice := redisClient.SMembers("ROOM:"+userMsg.Room)

	//用户uuid保存到一个go切片online
	online := roomSlice.Val()
	sendMsg,err2 := json.Marshal(rm)
	fmt.Println("broadcast")

	if err2 != nil {

	} else {
		//给所有用户发消息
		if len(online) != 0 {
			for _, na := range online {
				if na != "" {
					if err = ActiveClients[na].websocket.WriteMessage(websocket.TextMessage,sendMsg); err != nil {
						println("Could not send UsersList to ", "", err.Error())
					}
				}
			}
		}
	}
}

func NewUm() *UserMsg{
	return &UserMsg{}
}

