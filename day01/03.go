package main

import "fmt"

/* 在函数中使用管道之全局变量方式 */

var pipe chan int

func add1(a int, b int) {
	sum := a + b
	pipe <- sum
}

func main() {
	pipe = make(chan int, 3)
	go add1(39, 230)
	sum := <- pipe
	fmt.Println("sum =", sum)
}
