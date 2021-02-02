package main

import (
	"fmt"
	"net/http"
	"sync"
)

var ch chan string = make(chan string)

// Midler 中间件的操作
type Midler struct {
	Hands []http.HandlerFunc
	Run   int8 // 运行到第几个
	Rw    http.ResponseWriter
	Rq    *http.Request
	sync.RWMutex
}

func (h *Midler) next() {
	runFunc := h.Hands[int(h.Run)+1]
	runFunc(h.Rw, h.Rq)
	h.RLock()
	h.Run++
	h.RUnlock()
}

func (h *Midler) run() {
	for index, v := range h.Hands {
		if index >= int(h.Run) {
			if index == 0 {
				h.next()
			}
			v(h.Rw, h.Rq)
			h.RLock()
			h.Run++
			h.RUnlock()
		}
	}
}

// Midler 注册中间件 函数调用的过程
func (h *Midler) register(args ...http.HandlerFunc) {
	h.Hands = args
	fmt.Println(h.Hands)
}

func midelr(args ...http.HandlerFunc) http.HandlerFunc {
	md := &Midler{}
	return func(rw http.ResponseWriter, r *http.Request) {
		md.Rw = rw
		md.Rq = r
		md.register(args...)
		md.run()
	} // 函数执行的开始和函数执行的末尾实现
}

// next 方法实现的是一个递归调用的实现方式
func main() {
	http.HandleFunc("/", func(r http.ResponseWriter, q *http.Request) {
		http.Redirect(r, q, "/login", http.StatusMovedPermanently) // 实现页面中专的实现
		return
	})

	http.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hi i am login"))
	})

	http.HandleFunc("/func", midelr(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("one"))
	}, func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("two"))
	}, func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("three"))
	}))

	http.ListenAndServe(":8080", nil)
}

func consolea() {

}
