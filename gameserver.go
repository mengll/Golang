package server

import (
	"github.com/gorilla/websocket"
	"fmt"
	"time"
	"strconv"
	"math/rand"
	"github.com/go-redis/redis"
	"context"
	"encoding/json"
)

//用户的信息详情
type UserInfo struct {
	NickName 	string `json:"nick_name"`
	Avatar 		string `json:"avatar"`
	Gender 		string `json:"gender"`
}

//用户登录传递的数据
type UserDat struct {
	Cmd 		string `json:"cmd"`
	Uid 		string `json:"uid"`
	GameId      string `json:"game_id"`
	UserLimit   int    `json:"user_limit"`
	UserInfo
	Room 		string `json:"room"`
	RoomType    string `json:"room_type"`
}

//存放redis——header
const (
	CLIENT_LOGIN_KYE string         = "client_logined_game_key_%s"
	SUCESS_BACK int 				= 0
	FAILED_BACK int 				= 1
	RANDOM_USER  = "1"        //随机匹配
	PLAYER_REQ   = "2"        //玩家邀请
)

var (
	ActiveClients = make(map[string]map[string]ClientConn) //在线的用户的信息
    RedisClient *redis.Client
)

//创建游戏房间
func createRoom(gameid string) string {
	run_num  := time.Now().Second() //执行的时间戳
	rand_num := rand.Intn(999999)
	return fmt.Sprintf("%s_%d_%d",gameid,run_num,rand_num)
}

//从Redis的集合中移除数据
func delSet(key,val string){
	RedisClient.SRem(key,val)
}

//设置key
func setKey(k string ,v interface{}){
	RedisClient.Set(k,v,3600)
}


//添加到集合中
func addSet(key,val string){
	RedisClient.SAdd(key,val)
}

//判断是否存在
func hadSet(key,val string) bool{
	return RedisClient.SIsMember(key,val).Val()
}

//获取集合中的数量
func getSetNum(key string) int {
	checkNumTmp := RedisClient.SCard(key).Val()
	dt := strconv.FormatInt(checkNumTmp,10)
	dd,err := strconv.Atoi(dt)
	if err != nil{
		return 0
	}
	return dd
}

//数据类型定义
type ClientConn struct {
	websocket *websocket.Conn
}

type(
	//响应返回数据
	ResponseMsg struct {
		ErrorCode 	int 					`json:"error_code"`
		Data 		map[string]interface{}  `json:"data"`
		Msg 		string 					`json:"msg"`
	}

)

//发送广播
func clientBroadCast(c_room string,game_id string){
	c_members := RedisClient.SMembers(c_room)
	c_data :=  c_members.Val()
	//给房间内的所用玩家同步信息
	for _,v := range c_data{
		if _,ok := ActiveClients[game_id];ok {
			if con,oo := ActiveClients[game_id][v];oo{
				udat := RedisClient.Get("USER:"+v).Val()
				err := con.websocket.WriteJSON(udat)  //判断用户存在，则发送响应数据
				if err != nil{
					//println("发送用户信息失败:",data.Room,data.GameId)
					//
				}
			}
		}
	} //end for
}

//reday 列表中删除已开始游戏的对象
func delRedayMembers(reday,c_rooms string){
	cmbers := RedisClient.SMembers(c_rooms).Val()
	for _,v := range cmbers{

		delSet(reday,v)
	}
}

func WsInit(ws *websocket.Conn,udat *UserDat){
	uid 		 := udat.Uid
	game_id 	 := udat.GameId
	sockCli 	 := ClientConn{ws}
	rep 		 := ResponseMsg{}
	room_limit   := udat.UserLimit   //每个房间的人数限制

	//判断Redis连接情况
	redis_status := RedisClient.Ping()
	if _,err := redis_status.Result();err !=nil{
		rep.ErrorCode = FAILED_BACK
		rep.Msg = err.Error()
		ws.WriteJSON(rep)
	}

	switch udat.Cmd {
	case "login":
		println("start_login")
		login_key := fmt.Sprintf(CLIENT_LOGIN_KYE,udat.GameId)
		login_num := getSetNum(login_key)

		if ActiveClients[game_id] == nil{
			pk  := make(map[string]ClientConn)
			ActiveClients[game_id] = pk
		}
		ActiveClients[game_id][uid] = sockCli

		//保存用户登录信息
		addSet(login_key,uid)

		//保存用户的信息
		user_info := UserInfo{}
		user_info.Avatar   = udat.Avatar
		user_info.Gender   = udat.Gender
		user_info.NickName =udat.NickName

		//生成用户信息json串
		b, err := json.Marshal(user_info) //格式化当前的数据信息
		if err != nil {
			fmt.Println("Encoding User Faild")
		} else {

			//保存用户信息到redis
			RedisClient.Set("USER:"+udat.Uid, b, 0)
			println("保存用户信息到Redis--》")
			//初始化用户
			//initOnlineMsg(RedisClient,dat)
		}
		
		back_dat := make(map[string]interface{})
		back_dat["online_num"] = login_num + 1
		back_dat["game_id"] = game_id
		rep.ErrorCode = SUCESS_BACK
		rep.Data = back_dat
		rep.Msg = "login_sucess"
		ws.WriteJSON(rep)

	case "create_room":
		 new_room := createRoom(game_id)
		 limit_key := fmt.Sprintf("%s_limit",new_room)

		 //设置房间最大连接人数
		 setKey(limit_key,strconv.Itoa(room_limit))
		 addSet(new_room,"")
		 room_dat := make(map[string]interface{})
		 room_dat["room_id"] = new_room
		 rep.Msg = "create_room_sucess"
		 ws.WriteJSON(rep)

	case "join_room":
		 room := udat.Room
		 if room == ""{
		 	rep.ErrorCode = FAILED_BACK
		 	rep.Msg = "room not found"
		 	ws.WriteJSON(rep)
		 }

		 //当前房间的人数
		 room_num := getSetNum(room)
		 room_limit := RedisClient.Get(fmt.Sprintf("%s_limit",room)).Val()
		 num ,err := strconv.Atoi(room_limit)
		 if err != nil{
			 rep.ErrorCode = FAILED_BACK
			 rep.Msg = "room not found"
			 ws.WriteJSON(rep)
		 }

		 if num > room_num{

			 //加入成功
			 uid := udat.Uid
			 addSet(room,uid)
			 rep.ErrorCode = SUCESS_BACK

			 now_room_num := getSetNum(room)

			 if num == now_room_num{
			 	//start game
			 }

		 }else{

		 	//加入失败
		 	rep.ErrorCode = FAILED_BACK
		 	rep.Msg = "join_room_error"
		 	ws.WriteJSON(rep)
		 }

	case "search_match":
		room_limit := udat.UserLimit //游戏匹配的玩家的数量
		dd := []string{}
		gameReady := fmt.Sprintf("READY_RANDOM:%s",udat.GameId)      //所有准备的用户
		//当前房间的人数
		addSet(gameReady,udat.Uid)
		//todo 需要完善
		uid_channel := make(chan string,room_limit)
		//设置超时时间
		ctx,_ := context.WithTimeout(context.Background(),time.Second * 10)
		for i := 0;i<room_limit -1;i++ {

			go func(ctx context.Context) {
				//获取reday 列表中的数据数据总个数
				reday_num := getSetNum(gameReady)
				if reday_num >= room_limit {
				DEFR:
					rand_user := RedisClient.SPop(gameReady).Val()
					if rand_user != uid {
						uid_channel <- rand_user //随机生成要匹配的玩家的UID
					} else {
						goto DEFR
					}
				} else {
					ctx.Done() //结束当前
				}

			}(ctx)

		} //end for

		//获取其中的数据
		for i := 0;i<room_limit -1;i++ {
			select {
			case <-ctx.Done():
				//返回匹配超时
				close(uid_channel)
			case uid_c, ok := <-uid_channel:
				if ok == false {
					println("close")
					return
				}
				println(uid_c)
				dd =append(dd,uid_c)
				if len(dd) == room_limit{
					client_room  := createRoom(udat.GameId)
					for _,v := range dd{
						println(v) //广播当前的用户开始游戏
						addSet(client_room,v)
					}
					rep.ErrorCode = SUCESS_BACK
					rep.Data = map[string]interface{}{"cmd":"start"}
					rep.Msg = ""
					clientBroadCast(client_room,game_id) //广播通知当前的玩家，
					return  //结束当前的玩家
				}

			} //end select
		}//end for

	}
}

