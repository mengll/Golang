package main

import "fmt"
type Cx struct {
	Num int  // 当前的跑了第几次
	Left []int // 保存在左侧的数
	Right []int // 保存在右侧的数据
}
func main() {
	num := []int{1,4,2,3,5,9,10,11,24,14,34,13,45,17,19,40}
	
	back := compare(num)
	fmt.Println(back)
}

func compare(value []int)[]int{
	if len(value)==1 {
		return value
	}
	if len(value) == 0 {
		return nil
	}
	
	sl := Cx{}
	sl.Num =value[0]
	sl.Left = []int{}
	sl.Right =[]int{}

	for _,v := range value[1:] {
		if v > sl.Num {
			sl.Right = append(sl.Right,v)
		}else{
			sl.Left = append(sl.Left,v)
		}
	} 

	back := []int{}
	if len(sl.Right) >=1 {
		back = append(back,compare(sl.Right)...)
	}
	back = append(back,sl.Num)
	if len(sl.Left) >=1{
		back = append(back,compare(sl.Left)...)
	}
	return back
}

