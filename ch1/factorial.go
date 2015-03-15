package main

import "fmt"
import "math/big"

func main() {
	fmt.Println("Hello, playground")

	//n := big.NewInt(40)
	r := big.NewInt(15)

	fmt.Println(factorial(r))

}

func factorial(n *big.Int) (result *big.Int) {
	//fmt.Println("n = ", n)
	result = new(big.Int)

	switch n.Cmp(&big.Int{}) {
	case -1, 0:
		result.SetInt64(1)
	default:
		fmt.Println("n = ", n)
		result.Set(n)
		var one big.Int
		one.SetInt64(1)
		result.Mul(result, factorial(n.Sub(n, &one)))
	}
	return
}
