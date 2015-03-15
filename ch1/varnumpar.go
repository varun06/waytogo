package main

import "fmt"

func Min(a ...int) int {
	if len(a) == 0 {
		return 0
	}

	min := a[0]

	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
func main() {
	x := Min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	arr := []int{7, 9, 3, 5, 1}
	x = Min(arr...)
	fmt.Printf("The minimum of the array is: %d\n", x)
	x = Min(0)
	fmt.Printf("The minimum is: %d\n", x)
}
