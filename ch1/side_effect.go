package main

import "fmt"

var n int = 0

func Multiply(a, b int, reply *int) {
	*reply = a * b
}
func main() {
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply: ", *reply)
	fmt.Printf("Value of n is %d", n)
}
