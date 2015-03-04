package main

import "fmt"

func main() {
	// left shift is multiplied by 2^n
	// right shift is divided by 2^n
	n := 2
	m := 2048
	fmt.Printf("left shift is: %d\n", n<<3)
	fmt.Printf("right shift is: %d\n", m>>3)
}
3