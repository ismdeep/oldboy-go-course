package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	stop := int(math.Sqrt(float64((n))))
	for i := 2; i <= stop; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	for val := 100; val <= 200; val++ {
		if isPrime(val) {
			fmt.Println(val)
		}
	}
}