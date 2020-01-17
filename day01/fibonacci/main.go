package main

import "fmt"

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n - 1) + fibonacci(n - 2)
}

func fibonacciCalc(n int, pipe chan int) {
	pipe <- fibonacci(n)
}

func main() {
	pipe := make(chan int, 50)
	for i := 1; i <= 100; i++ {
		go fibonacciCalc(40, pipe)
	}
	for i := 1; i <= 100; i++ {
		fmt.Println(<- pipe)
	}
}