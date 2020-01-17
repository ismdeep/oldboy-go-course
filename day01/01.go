package main

import (
	"fmt"
	"time"
)

func testGoroutine(a int, b int) {
	sum := a + b
	fmt.Println("testGoroutine()", sum)
}

func testPrint(a int) {
	fmt.Println("testPrint()", a)
}

func main() {
	for i := 0; i < 100; i++ {
		go testPrint(i)
	}
	go testGoroutine(300, 600)
	time.Sleep(time.Second)
}
