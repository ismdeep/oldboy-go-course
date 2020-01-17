package main

import "fmt"

func isOk(n int) bool {
	tmp := n
	a := tmp % 10
	tmp /= 10
	b := tmp % 10
	tmp /= 10
	c := tmp
	if a * a * a + b * b * b + c * c * c == n {
		return true
	}
	return false
}

func main() {
	for val := 100; val <= 999; val++ {
		if isOk(val) {
			fmt.Println(val)
		}
	}
}