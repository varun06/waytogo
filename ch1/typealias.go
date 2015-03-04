package main

import "fmt"

type TZ int
type rope string

func main() {
	var a, b TZ = 3, 4
	e := a + b
	var c, d rope = "hello", "world"
	fmt.Printf("c has the value: %d\n", e)
	fmt.Printf("%s\n", c)
	fmt.Printf("%s\n", d)
}
