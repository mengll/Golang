# Golang
### 切片清空  	chans = chans[:0]
### 数据格式
数据都是补码的形式存在
// 大转小获取，小类型的后几位， 获取数据补码 负数的时候 =》 反码 + 1 正数的时候 》 正常的补码
### interface 类型断言 interface{}(a).(string) interface{}类型转化下，在进行断言
# 通过make创建的切片对外是不可见的，只能通过slice的方式处理数据 make的方式创建数据的访问地址（这很重要）

## 对于一个数据类型的断言
```
var jk = "12"

	switch interface{}(jk).(type) {
	case string:
		fmt.Println("This is 啊string ")
	}


type ty struct {
   Name string `json:"name"`
}
```
部分被称为标记tag  表示的是json 解码的元数据 用于创建ty数据类型原值的切片 使用这个函数读取数据
```

```


	func SetCookie(k, v string, t int, w *http.ResponseWriter) {
		COOKIE_MAX_MAX_AGE := t // 单位：秒。
		maxAge := int(COOKIE_MAX_MAX_AGE)

		uid_cookie := &http.Cookie{
			Name:     k,
			Value:    v,
			Path:     "/",
			HttpOnly: false,
			MaxAge:   maxAge}
		http.SetCookie(*w, uid_cookie) //
	}

	//获取Key
	func GetCookie(k string, r *http.Request) string {
		cokcont, _ := r.Cookie(k)
		return cokcont.Value //获取当前缓存的K
	}
	
	//将str转换为时间格式
	func StrToTime(st string) time.Time {
		t, _ := time.ParseInLocation(f_datetime, st, time.Local) //时间戳转化
		return t
	}
	
	postdate := StrToTime("2017-09-11 00:00:00")
	
	// 渲染页面并输出
	func renderHTML(w http.ResponseWriter, file string, data interface{}) {
		// 获取页面内容
		t, err := template.New(file).ParseFiles("views/" + file)
		checkErr(err)
		// 将页面渲染后反馈给客户端
		t.Execute(w, data)
	}
	
	func checkErr(err error) {
		if err != nil {
			log.Println(err)
		}
	}
       //检查当前程序的竞态
       go run -race  main.go
	// 随机延时调用
	now_hour := time.Now().Hour()
	af_time := 10
	rand_int := rand.Intn(time.Now().Nanosecond())
	ra, _ := strconv.Atoi(strconv.Itoa(rand_int)[:1])

	if now_hour < 6 {
		af_time = 10 + ra
	} else {
		af_time = 3 + ra
	}

	select {

	case <-time.After(time.Second * time.Duration(af_time)):
	
	func Log(txt interface{}) {
		pc, file, line, ok := runtime.Caller(1)
		f := runtime.FuncForPC(pc)
		fmt.Println("【pkgame】", fmt.Sprintf("func = %s,file = %s,line = %d,ok = %v ,val = %v", f.Name(), file, line, ok, txt))
	}

	
```

```
动态设置获取的值

type Gamea struct {
	Name string `json:"name,username"`
	Age int `json:"age"`
}

func main(){
	nam := `{"username":"fja","age":12}`
	Gm := Gamea{}
	//bb,_ := json.MarshalIndent(Gm,nam,"\t")
	json.Unmarshal([]byte(nam),&Gm)
	fmt.Println(Gm.Name,Gm.Age)

	t := reflect.TypeOf(Gm)
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

		}

		}
	}

	fmt.Println(Gm.Age)
	fmt.Println(Gm.Name)
}

```


```
执行相关的单元测试命令
go test -v main_test.go main.go  第一个参数测试文件，第二个是被测试的文件

2对单个方法进行测试
go test -v -test.run TestFunc   TestFunc 是想要执行的测试的测试函数

```

```
go变种操作
import (
    "fmt"
    "sync"
    "time"
)

type WaitGroupWrapper struct {
    sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func(argvs ...interface{}), argvs ...interface{}) {
    w.Add(1)
    go func() {
        cb(argvs...)
        w.Done()
    }()
}

###获取字符串的真是长度 而不是字节长度
使用普通的len()获取的是字节的长度 中文字符占据了三个字节 要想获取真正的字符长度那就
utf8.RuneCountInString(name)

base64的数据的长度必须是满足 4的倍数，如果不满足需要补上 = 缺少几个就需要补几个

```
```
### 程序信号
		// 信号终止程序退出
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT) // 进程终止（软件终止），中断进程 ，杀死进程，中止信号，建立Core文件终止进程
	go func() {
		s := <-ch
		log.Printf("receive signal '%v'", s)
		os.Exit(1)
	}()

```

### 生成随机数的操作
```	
rd := rand.New(rand.NewSource(time.Now().Unix()))
fmt.Println(rd.Intn(100))
```	
### vscode 的go配置
```
"go.gopath":"${workspaceRoot}:/Users/Young/Desktop/go", // 当前工作空间${wordspaceRoot}加上系统 GOPATH 目录
"go.goroot": "/usr/local/Cellar/go/1.12/libexec", // go 的安装目录
"go.formatOnSave": true, //在保存代码时自动格式化代码
"go.formatTool": "goimports", //使用 goimports 工具进行代码格式化，或者使用 goreturns 和 gofmt
"go.buildOnSave": true, //在保存代码时自动编译代码
"go.lintOnSave": true, //在保存代码时自动检查代码可以优化的地方，并给出建议
"go.vetOnSave": false, //在保存代码时自动检查潜在的错误
"go.coverOnSave": false, //在保存代码时执行测试，并显示测试覆盖率
"go.useCodeSnippetsOnFunctionSuggest": true, //使用代码片段作为提示
"go.gocodeAutoBuild": false //代码自动编译构建

链接：https://juejin.im/post/5c7c8fdf518825763c6d9cb
```



### 配置访问的代理
export GO111MODULE=on

export GOPROXY=https://goproxy.io

### 高质量文章
 go调度 https://www.cnblogs.com/zkweb/category/1108329.html
 
 https://studygolang.com/articles/11627
 
 go内存  https://studygolang.com/articles/21033
 
 博客  https://xargin.com/

slice 详解 https://www.cppentry.com/bencandy.php?fid=78&aid=216871&page=1

优秀资源汇总
https://shockerli.net/post/go-awesome/

### dlv 的安装
```
-- 1.16 以下的安装方式
$ git clone https://github.com/go-delve/delve
$ cd delve
$ go install github.com/go-delve/delve/cmd/dlv

--- 1.16 或更高的版本
# Install the latest release:
$ go install github.com/go-delve/delve/cmd/dlv@latest

# Install at tree head:
$ go install github.com/go-delve/delve/cmd/dlv@master

# Install at a specific version or pseudo-version:
$ go install github.com/go-delve/delve/cmd/dlv@v1.7.3
$ go install github.com/go-delve/delve/cmd/dlv@v1.7.4-0.20211208103735-2f13672765fe

```






