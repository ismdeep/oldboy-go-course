package main

import "fmt"

/* 在函数中使用管道之传入参数方式 */

func add2(a int, b int, pipe1 chan int) {
	sum := a + b
	pipe1 <- sum
}

func main() {
	var pipe chan int
	pipe = make(chan int, 3)
	add2(20, 34, pipe)
	sum := <- pipe
	fmt.Println("sum =", sum)
}