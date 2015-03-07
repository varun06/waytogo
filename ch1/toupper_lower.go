package main

import (
	"fmt"
	"strings"
)

func main() {
	var orig string = "Hey, how are you Guys?"
	var lower string
	var upper string

	fmt.Printf("The original string is: %s\n", orig)
	lower = strings.ToLower(orig)
	fmt.Printf("The lowercase string is: %s\n", lower)
	upper = strings.ToUpper(orig)
	fmt.Printf("The upperacse string is: %s\n", upper)
}
