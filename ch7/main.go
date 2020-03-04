package main

import (
	"flag"
	"fmt"
	"time"

	"gopl.io/ch7/bytecounter"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	var c bytecounter.Bytecounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
