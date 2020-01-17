package main

import (
	"fmt"
	"time"
)

func main() {
	select {
	case <-time.After(time.Second):
		fmt.Println("after")
	}
}