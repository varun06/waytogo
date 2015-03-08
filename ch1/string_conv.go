package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "666"
	var an int
	var newS string
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	//string to integer
	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)
	an += 5
	//integer to string
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
