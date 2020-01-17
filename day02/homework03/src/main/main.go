package main

import "fmt"

func factorial(n uint64) uint64 {
	var result uint64 = 1
	for i := uint64(1); i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Println(factorial(10))
}
