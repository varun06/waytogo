package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := []byte("asSASA ddd dsjkdsjs ^^&& dk")
	fmt.Println("bytes =", len(s))
	fmt.Println("runes =", utf8.RuneCount(s))
}
