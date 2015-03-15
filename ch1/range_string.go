package main

import "fmt"

func main() {
	str := "Go is a beautiful language"
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c\n", pos, char)
	}
	fmt.Println()
	str2 := "Chinese: 围棋是一种美丽的语言"
	for pos, char := range str2 {
		fmt.Printf("Character on position %d is: %c\n", pos, char)
	}
	fmt.Println()
}
