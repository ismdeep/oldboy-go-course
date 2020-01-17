package main

import (
	"fmt"
	"math/big"
)

func main() {
	var (
		aStr string
		bStr string
	)
	fmt.Printf("Input big int a: ")
	fmt.Scanf("%s", &aStr)
	fmt.Printf("Input big int b: ")
	fmt.Scanf("%s", &bStr)
	big1, _ := new(big.Int).SetString(aStr, 10)
	big2, _ := new(big.Int).SetString(bStr, 10)
	bigSum := new(big.Int).Add(big1, big2)
	fmt.Println(bigSum.String())
}