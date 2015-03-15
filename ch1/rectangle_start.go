package main

import "fmt"

func main() {
	var w, h = 20, 10
	for i := 0; i <= h; i++ {
		for j := 0; j < w; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
