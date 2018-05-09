# Golang
###切片清空  	chans = chans[:0]
###interface 类型断言 interface{}(a).(string) interface{}类型转化下，在进行断言

##对于一个数据类型的断言
var jk = "12"

	switch interface{}(jk).(type) {
	case string:
		fmt.Println("This is 啊string ")
	}

``` 

type ty struct {
   Name string `json:"name"`
}
`` 部分被称为标记tag  表示的是json 解码的元数据 用于创建ty数据类型原值的切片 使用这个函数读取数据
```

```

//创建
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
