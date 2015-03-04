package main

import "fmt"

func main() {
	var a int
	var b int32
	a = 15

	b = int32(a) + 5
	b = b + 5
	fmt.Printf("%05d %03d\n", a, b)
}
