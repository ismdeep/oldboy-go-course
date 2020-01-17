package main

import "fmt"

func main() {
	// 打印字符串
	fmt.Println("Hello World.")
	// 打印二进制
	fmt.Printf("%b\n", 10)
	// 打印十进制
	fmt.Printf("%d\n", 23409)
	// 打印十六进制(小写)
	fmt.Printf("%x\n", 239847)
	// 打印十六进制(大写)
	fmt.Printf("%X\n", 239847)
	// 打印带0X的十六进制
	fmt.Printf("%#X\n", 239847)
	// 打印浮点数
	fmt.Printf("%f\n", 23490.3)
}