package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) uint64 {
	if n <= 2 {
		return uint64(1)
	}
	return fibonacci(n - 1) + fibonacci(n - 2)
}

func main() {
	start_time := time.Now()
	fmt.Println(fibonacci(44))
	end_time := time.Now()
	fmt.Printf("cost: %d us\n", (end_time.UnixNano() - start_time.UnixNano()) / 1000)
	fmt.Println(start_time.Unix())
}