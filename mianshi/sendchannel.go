package main

import (
	"fmt"
	"sync"
	"time"
)

var wt sync.WaitGroup

func main() {
	ch_nume := make(chan bool)
	ch_char := make(chan bool)
	wt.Add(2)
	go func() {
		ch_nume <- true
	}()
	go func(wt *sync.WaitGroup) {
		ste := 0
		for range ch_nume {
			ste++
			fmt.Println(ste)
			ste++
			fmt.Println(ste)
			time.Sleep(time.Second)
			if ste == 28 {
				wt.Done()
				close(ch_nume)
				close(ch_char)
				return
			} else {
				ch_char <- true
			}
		}
	}(&wt)

	go func(wt *sync.WaitGroup) {
		step := 0
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for range ch_char {
			fmt.Println(str[step : step+1])
			step++
			fmt.Println(str[step : step+1])
			step++
			ch_nume <- true
		}
		wt.Done()
	}(&wt)
	wt.Wait()
}
