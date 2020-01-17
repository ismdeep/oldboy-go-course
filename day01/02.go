package main

import "fmt"

func testPipe() {
	// 创建一个管道
	pipe := make(chan int, 3)
	// 将数据放入管道中
	pipe <- 1
	pipe <- 2
	// 查看管道中数据量
	fmt.Println("pipe size:", len(pipe))
	pipe <- 3
	fmt.Println("pipe size:", len(pipe))
	// 从管道中取出数据
	val := <- pipe
	fmt.Println("val:", val)
}

func main() {
	testPipe()
}