package main

import "fmt"

func PrintSymbolTrangle(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("A")
		}
		fmt.Printf("\n")
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	PrintSymbolTrangle(n)
}
