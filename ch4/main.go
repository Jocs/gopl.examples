package main

import (
	"fmt"

	"gopl.io/ch4/append"
)

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = append.AppendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
