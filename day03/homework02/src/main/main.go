package main

import "fmt"

func genFactors(n int) []int {
	var a []int
	for i := 1; i < n; i++ {
		if n % i == 0 {
			a = append(a, i)
		}
	}
	return a
}

func sumSlice(a []int) int {
	var sum int = 0
	for _, item := range a {
		sum += item
	}
	return sum
}

func isPerfectNumber(n int) bool {
	if n == sumSlice(genFactors(n)) {
		return true
	} else {
		return false
	}
}

func main() {
	for val := 1; val <= 1000; val++ {
		if isPerfectNumber(val) {
			fmt.Println(val)
		}
	}
}