package main

import "fmt"

func main() {
	sum1, prod1, diff1 := SumProdDiff(5, 4)
	sum2, prod2, diff2 := SumProdDiff2(8, 6)
	fmt.Printf("Sum is: %d\n prod is: %d\n diff is: %d\n", sum1, prod1, diff1)
	fmt.Printf("Sum2 is: %d\n prod2 is: %d\n diff2 is: %d\n", sum2, prod2, diff2)
}

func SumProdDiff(num1, num2 int) (int, int, int) {
	sum := num1 + num2
	prod := num1 * num2
	diff := num1 - num2
	return sum, prod, diff
}

func SumProdDiff2(num3, num4 int) (sum2, prod2, diff2 int) {
	sum2 = num3 + num4
	prod2 = num3 * num4
	diff2 = num3 - num4
	return
}
