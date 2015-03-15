package main

import "fmt"

func main() {
	n := 5
	//
	DoSomething(&n)
	DoSomething2(n)
}

func DoSomething(a *int) {
	b := a
	fmt.Println(b)
}

func DoSomething2(a int) {
	b := &a
	fmt.Println(b)
}
