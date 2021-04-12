package main

import "fmt"

func main() {
	i := 1
	for foo("A", i); foo("B", i) && i < 2; foo("C", i) {
		i++
		foo("D", i)
	}
}

func foo(c string, i int) bool {
	fmt.Println(c, i)
	return true
}

// A 1
// B 1
// D 2
// C 2
// B 2
