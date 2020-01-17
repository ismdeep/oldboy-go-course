package main

import (
	"fmt"
	"math/big"
)

func genFibLst(n int) (a []string) {
	n1 := big.NewInt(1)
	n2 := big.NewInt(1)
	for n > 0 {
		a = append(a, n1.String())
		tmp := n1.Add(n1, n2)
		n1 = n2
		n2 = tmp
		n--
	}
	return a
}

func main() {
	lst := genFibLst(1000)
	for index, val := range lst {
		fmt.Println(index + 1, val)
	}
}