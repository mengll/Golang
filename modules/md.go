package main

import (
	"fmt"
	"sync"
	"testing"
)

type Dat struct {
	Name string `json:"name"`
	Muid string `json:"muid"`
	Pid  int32  `json:"pid"`
}

type Fn func(*Middle, *Dat)

type Middle struct {
	M       sync.Mutex
	Handles []Fn // 数据格式的初始化
	Run     int32
	Data    *Dat
}

type Hjd interface {
	Next()
}

func (m *Middle) Next() {
	if len(m.Handles) > int(m.Run) {
		m.M.Lock()
		f := m.Handles[m.Run+1]
		m.Run += 1
		fmt.Println(m.Run, "---->")
		m.M.Unlock()
		f(m, m.Data)
	}
}

func (m *Middle) Start() {
	for i := int(m.Run); i < len(m.Handles); {
		f := m.Handles[m.Run]
		f(m, m.Data)
		m.M.Lock()
		m.Run++
		m.M.Unlock()
		i = int(m.Run)
	}
}

func (m *Middle) Add(f Fn) {
	m.Handles = append(m.Handles, f)
}

// 初始化 midleware
func NewMiddle() *Middle {
	return &Middle{Run: 0, Handles: []Fn{}, Data: &Dat{Name: "Shandong", Muid: "the word is big", Pid: 1}}
}

func one(m *Middle, dat *Dat) {
	dat.Pid = 1
	fmt.Println("one")
}
func two(m *Middle, dat *Dat) {
	fmt.Println("two")
	dat.Pid = 2
	m.Next()
	fmt.Println("two after")
}

func three(m *Middle, dat *Dat) {
	fmt.Println("three")
	dat.Pid = 3
}

func four(m *Middle, dat *Dat) {
	fmt.Println("four")
	dat.Pid = 4
}

func Test_main(t *testing.T) {
	fmt.Println("test")
	m := NewMiddle()
	m.Add(one)
	m.Add(two)
	m.Add(three)
	m.Add(four)
	m.Start()
	fmt.Println(m.Data)
}

func main() {
	m := NewMiddle()
	m.Add(one)
	m.Add(two)
	m.Add(three)
	m.Add(four)
	m.Start()
}
