package v1

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"smallgamepk.qcwanwan.com/libs/db"
)

type Roomer interface {
	JoinRoom(room, uid string) (bool, error) //加入房间
	OutRoom(room, uid string) (bool, error)  //离开房间
	RoomLimit(game_id string) (int, error)   //房间人数限制
	IsFull(room string) (bool, error)        //房间是否满员
	ExistsRoom(room string) (bool, error)    //创建的房间是否还在
	NumRoom(room string) (int, error)        //当前房间的人数
	CloseRoom(room string) error             //关闭当前的房间

	CreateRoom(game_id, uid, share string) (string, error) // 返回当前的房间名称和错误信息
	ClearnRoom(room string) error
}

type RoomManager struct{}

var Pgredis db.GsRedis = db.NewRedis()

func init() {
	Pgredis.Connect()
}

/*
@return string the create room name
@return error  error message
@author mengll
*/

func (self *RoomManager) CreateRoom(game_id, uid, share string) (string, error) {

	run_num := time.Now().Unix() //执行的时间戳
	rand_num := rand.Intn(999999)
	new_room := ""

	if share == ROOM_TYPE_SHARE {
		new_room = fmt.Sprintf("ROOM:%s_%d_%d_%s", game_id, run_num, rand_num, share)
	} else {
		new_room = fmt.Sprintf("ROOM:%s_%d_%d", game_id, run_num, rand_num)
	}

	if game_id == "" || len(game_id) > 100 {
		return "", errors.New("game_id is nil or out of range")
	}

	limit_key := fmt.Sprintf("%s_limit", new_room)

	//设置房间最大连接人数
	room_limt := 2
	room_limt, _ = self.RoomLimit(game_id)

	Pgredis.SetKey(limit_key, room_limt)
	Pgredis.AddSet(new_room, uid)
	//Pgredis.Expire(new_room,time.Second * 3600) //设置房间的失效时间
	return new_room, nil
}

/*
@return error message
*/
func (self *RoomManager) CloseRoom(room string) error {
	if strings.HasPrefix(room, "ROOM") {
		Pgredis.DelKey(room)
	}

	return errors.New("close room error room type error")
}

/*
@return int the num of the room
*/

func (self *RoomManager) NumRoom(room string) (int, error) {

	if strings.HasPrefix(room, "ROOM:") {
		return Pgredis.GetSetNum(room), nil
	}

	return 0, nil
}

/*
@bool 当前房间时候还存在
*/
func (self *RoomManager) ExistsRoom(room string) (bool, error) {

	if strings.HasPrefix(room, "ROOM:") {
		return Pgredis.EXISTS(room)
	}

	return false, nil
}

/*
@return bool the room num is more than room limit
*/
func (self *RoomManager) IsFull(room string) (bool, error) {

	if strings.HasPrefix(room, "ROOM:") {

		room_limit, _ := self.RoomLimit("")

		if Pgredis.GetSetNum(room) >= room_limit {
			return true, nil //当前房间已满
		}

		return false, nil
	}
	return false, errors.New("room message error")
}

/*
@return the max user in the room
*/
func (self *RoomManager) RoomLimit(game_id string) (int, error) {
	return 2, nil
}

/*
@return bool is sucess out of the room
*/
func (self *RoomManager) OutRoom(room, uid string) (bool, error) {
	if strings.HasPrefix(room, "ROOM:") {
		if len(uid) == 0 {
			return false, errors.New("uid is nil")
		}
		if Pgredis.GetSetNum(room) == 0 {
			Pgredis.DelKey(room)
		}

		err := Pgredis.DelSet(room, uid)
		if err != nil {
			return false, err
		}

		return true, nil
	}
	return false, errors.New("room message error") //
}

/*
@bool 加入房间
#todo 房间判断后期优化
*/
func (self *RoomManager) JoinRoom(room, uid string) (bool, error) {

	if strings.HasPrefix(room, "ROOM:") {

		if len(uid) == 0 {
			return false, errors.New("【进入房间缺少用户id】")
		}

		if Pgredis.HadSet(room, uid) {
			return false, errors.New("【用户已经在房间中】")
		}

		room_exists, room_err := Pgredis.EXISTS(room)
		if room_err != nil || !room_exists {
			return false, errors.New("【您要加入的房间已经关闭】")
		}

		room_num := Pgredis.GetSetNum(room)
		room_limit := Pgredis.GetKey(fmt.Sprintf("%s_limit", room))
		num, err := strconv.Atoi(room_limit)

		if err != nil {
			fmt.Println("【转换房间限制失败】")
			num = 2
		}

		if num > room_num {
			Pgredis.AddSet(room, uid)
			now_room_num := Pgredis.GetSetNum(room)
			if now_room_num == room_num {
				return true, nil
			}

		} else {
			return true, errors.New("【房间已满】")
		}
	}

	return false, errors.New("room message error")
}

//清除房间
func (self *RoomManager) ClearnRoom(room string) error {
	Pf_redis.DelKey(room)
	return nil
}

//返回room manager
func NewRoomManager() Roomer {
	return new(RoomManager)
}
