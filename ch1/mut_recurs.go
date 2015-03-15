package main

import "fmt"

func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}

func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}

func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}

func main() {
	fmt.Printf("%d is even: is %t\n", 16, even(16))
	fmt.Printf("%d is odd: is %t\n", 17, odd(17))
	fmt.Printf("%d is odd: is %t\n", 16, odd(16))
}
