package server

import (
	"github.com/go-redis/redis"
	"strconv"
)

type(
	GsRedis interface {
		Connect()
		getSetNum(string) int
		hadSet(key,val string) bool
		addSet(key,val string) error
		setKey(k string ,v interface{}) error
		delSet(key,val string) error
		GetKey(k string)string
		SPop(k string) string
		DelKey(k string) error
		SMembers(k string)([]string,error)
	}

	GsRedisManage struct {
		Address 	string
		Passworld   string
		DB      	int
		RS     		*redis.Client
	}
)

func (this *GsRedisManage) Connect(){
	   this.RS= redis.NewClient(&redis.Options{
		Addr:     this.Address,
		Password: this.Passworld, // 设置Redis的链接的链接方法
		DB:       this.DB,  // use default DB
	   })
}

//从Redis的集合中移除数据
func (this *GsRedisManage)delSet(key,val string) error{
	return this.RS.SRem(key,val).Err()
}

//设置key
func (this *GsRedisManage)setKey(k string ,v interface{})error{
	return this.RS.Set(k,v,0).Err()
}

//添加到集合中
func (this *GsRedisManage)addSet(key,val string) error{
	return this.RS.SAdd(key,val).Err()
}

//判断是否存在
func (this *GsRedisManage)hadSet(key,val string) bool{
	return this.RS.SIsMember(key,val).Val()
}

//获取数据
func (this *GsRedisManage) GetKey(k string) string {
	return this.RS.Get(k).Val()
}

//随机删除并返回
func(this *GsRedisManage)SPop(k string) string{
	return this.RS.SPop(k).Val()
}

//删除键
func(this *GsRedisManage)DelKey(k string) error{
	return this.RS.Del(k).Err()
}

//获取集合中的内容
func(this *GsRedisManage)SMembers(k string)([]string,error){
	set_val := this.RS.SMembers(k)
	return set_val.Val(),set_val.Err()
}

//获取集合中的数量
func (this *GsRedisManage)getSetNum(key string) int {
	checkNumTmp := this.RS.SCard(key).Val()
	dt := strconv.FormatInt(checkNumTmp,10)
	dd,err := strconv.Atoi(dt)
	if err != nil{
		return 0
	}
	return dd
}

//生成当前redis 对象
func NewRedis() GsRedis{
	redis_client :=  new(GsRedisManage)
	redis_client.Address   = "192.168.1.246:6379"
	redis_client.Passworld =  ""
	redis_client.DB = 1
	return redis_client
}

