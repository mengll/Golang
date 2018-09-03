package main

import (
	"fmt"
	"reflect"
	"encoding/json"
	"strings"
	"os"
	"runtime/trace"
)
//主流的数据传输格式，序列化格式化的，字符串，反序列化的过程中有一个概念
type Gh interface {
	show() string
}


type dejh struct {
	Name string
}

func (self *dejh) show() string{
	fmt.Println("This is show func ")
	return ""
}

var jk = "fk"
func mainshow(){
	djk := []string{"23","356"}
	fmt.Println(djk)
	djk = djk[:0] //清空当前的对象
	fmt.Println(djk)//字符串的输出偶成写成完整的内容的实例的体现的过程

	var sjd Gh = new(dejh)
	sjd.show() //调用了show function 操作过程中的方法 返回相关的内容

	switch interface{}(jk).(type) {
		case string:
			fmt.Println("This is 啊string ")
	}
	
}
type Gamer interface{
	show()
}

type Game struct {
	Name string
}

func (self *Game)Show(){
	fmt.Println("This is the best")
}

func Fshow(a int){
	fmt.Println("This is show func ")
}

//创建文件
func mainjk(){
	 a := new(Game)
	 a.Name = "wendan"
	fu := reflect.ValueOf(a)
	mv := fu.MethodByName("Show")
	mv.Call([]reflect.Value{}) // 调取相关的值

	//创建分享的字段
	t := reflect.TypeOf(a)
	fmt.Println(t.Elem().Kind()) //定义对象额类型，获取相关的设置

	for i:= 0;i < t.NumMethod();i++ {
		fmt.Println(t.Method(i).Name)
	}

	if t.Elem().Kind() == reflect.Struct {
		fmt.Println("This is struct")
		for j:=0;j<t.Elem().NumField();j++{ //广告公司的传说的事情
			fmt.Println("222")
			fmt.Println(t.Elem().Field(j).Tag.Get("Name"))
		}

		if field, ok := t.Elem().FieldByNameFunc(func(s string) bool {
			fmt.Println(s)
			return len(s) > 3
		}); ok {
			fmt.Println("FieldByNameFunc    :", field.Name)
		}

	}

	var hj Game

	fmt.Println(hj)

}

type Gamea struct {
	Name string `json:"name,username"`
	Age int `json:"age"`
}

func mainodk(){
	f,err := os.Create("trace.out")
	if err != nil{
		panic(err)
	}
	defer f.Close()
	err_trace := trace.Start(f)
	if err_trace != nil{
		panic(err_trace)
	}
	nam := `{"username":"fja","age":12}`
	Gm := Gamea{}
	//bb,_ := json.MarshalIndent(Gm,nam,"\t")
	json.Unmarshal([]byte(nam),&Gm)  // 解析数据类型的内容
	fmt.Println(Gm.Name,Gm.Age)

	t := reflect.TypeOf(Gm)  //数据类型 kind类别
    vv := reflect.ValueOf(&Gm)

	for i:=0;i<t.NumField();i++{
		names := t.Field(i).Tag.Get("json")
		for _,v := range strings.Split(names,","){
			fmt.Println(v)
				//println("Ths is age ",t.Field(i).Name)
		jkk := vv.Elem().FieldByName(t.Field(i).Name)
		switch t.Field(i).Type.Kind() {
		case reflect.Int:
			jkk.SetInt(19)
		case reflect.String:
			jkk.SetString("mlove")

		} //类型翻转

		}
	}

	fmt.Println(Gm.Age)
	fmt.Println(Gm.Name)

	defer trace.Stop()
}

 type WD struct {
  	Name string
  	Age  int
 }

 func (wd *WD)Fuck(){
 	fmt.Println("pil type")
 	wd.Name = "wendan"
 }

 func (self WD)Gook(){
 	fmt.Println("This is Gooka")
 	self.Name = "fjka"
 }


func main(){
	var Gm Game
	if Gm == (Game{}){
		fmt.Println("This is empty")
	}
	 var x,y int //整天的复制操作
	 x,y = 2,3
	 fmt.Println(x,y)
	//判断当前的干系

	kl:= WD{}
	kl.Fuck()
    kl.Gook()

    hj := new(WD)
    hj.Gook()
    mp := make(map[string]func(a int))
    mp["fk"] = Fshow
    mp["fk"](2)
	fmt.Println(kl)
}
