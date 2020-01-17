package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.RWMutex
	fmt.Println(mu)
}
