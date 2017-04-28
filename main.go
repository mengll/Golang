// TestTool project main.go
package main

//"TestTool/utest"
import (
	"TestTool/controllers"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

// 渲染页面并输出
func renderHTML(w http.ResponseWriter, file string, data interface{}) {
	// 获取页面内容
	t, err := template.New(file).ParseFiles("views/" + file)
	if err != nil {
		fmt.Println("完成规则")
	}

	checkErr(err)
	// 将页面渲染后反馈给客户端
	t.Execute(w, data)
}

// 处理用户提交的数据
func page(w http.ResponseWriter, r *http.Request) {
	// 我们规定必须通过 POST 提交数据
	if r.Method == "POST" {
		controllers.StrData(w, r)
		return

	} else {
		// 如果不是通过 POST 提交的数据，则将页面重定向到主页
		renderHTML(w, "redirect.html", "/")
	}
}

// 处理主页请求
func index(w http.ResponseWriter, r *http.Request) {
	// 渲染页面并输出
	renderHTML(w, "login.html", "no data")
}

func main() {
	//fmt.Println(utest.HttpDO("POST", "/login", "app_version=12asd123"))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) //设置静态资源的访问路径
	//创建测试脚本的实现
	http.HandleFunc("/", index) // 设置当前的页面的访问的限制的控制的处理

	http.HandleFunc("/page", page) // 设置访问的路由

	//创建访问的路由
	//http.HandleFunc("/login", controllers.UserLogin)
	http.HandleFunc("/login", controllers.Newlogin)

	//创建新的请求
	http.HandleFunc("/newrequest", controllers.NewRequest)

	//测试访问控制
	http.HandleFunc("/newsign", controllers.NewSign)

	//获取当前的渠道信息
	http.HandleFunc("/getuserinfo", controllers.GetUserInfo)

	//解析数据的方式获取请求的具体的连接方式和，请求地址
	http.HandleFunc("/data", controllers.StrData)
	//登录后请求地址
	http.HandleFunc("/homepage", controllers.HomePage)

	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
