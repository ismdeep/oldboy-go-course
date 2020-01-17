package main

import "fmt"

/* 函数多返回值 */
func calc(a int, b int) (int, int) {
	sum := a + b
	avg := sum / 2
	return sum, avg
}

func main() {
	sum, avg := calc(2, 4)
	fmt.Println("sum, avg =", sum, avg)
}

