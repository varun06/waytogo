package main

import "fmt"

func main() {
	s := "good bye"
	var p *string = &s
	// *p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p)
	fmt.Printf("Here is the string *p: %s\n", *p)
	fmt.Printf("Here is the string s: %s\n", s)
	// * is used for creating pointers, for dereferenceing a pointer
	// for multiplication
}
