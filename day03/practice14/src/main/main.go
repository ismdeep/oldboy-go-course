package main

import "fmt"

func sum(a int, arg...int) int {
	ans := a
	for _, val := range arg {
		ans += val
	}
	return ans
}

func main() {
	fmt.Println(sum(1,2, 3))
}