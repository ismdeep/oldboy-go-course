package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	stop := int(math.Sqrt(float64(n)))
	for i := 2; i <= stop; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func calc(taskChan chan int, resultChan chan int, exitChan chan bool) {
	for v := range taskChan {
		if isPrime(v) {
			resultChan <- v
		}
	}
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)
	go func() {
		for i := 0; i < 100000; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	for i := 0; i < 8; i++ {
		go calc(intChan, resultChan, exitChan)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
			fmt.Println("Goroutine", i, "exited.")
		}
		close(resultChan)
	}()

	for prime := range resultChan {
		fmt.Println(prime)
	}
}
