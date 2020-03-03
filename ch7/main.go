package main

import (
	"fmt"

	"gopl.io/ch7/bytecounter"
)

func main() {
	var c bytecounter.Bytecounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
