package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	buff := bufio.NewReader(os.Stdin)
	str, _ = buff.ReadString('\n')
	str = str[0:len(str)-1]
	digitCnt := 0
	alphabetCnt := 0
	whitespaceCnt := 0
	otherCnt := 0
	for _, item := range str {
		if ' ' == item {
			whitespaceCnt++
		} else if '0' <= item && item <= '9' {
			digitCnt++
		} else if ('a' <= item && item <= 'z') || ('A' <= item && item <= 'Z') {
			alphabetCnt++
		} else {
			otherCnt++
		}
	}
	fmt.Printf("Alphabet Count: %d\n", alphabetCnt)
	fmt.Printf("Whitespace Count: %d\n", whitespaceCnt)
	fmt.Printf("Digital Count: %d\n", digitCnt)
	fmt.Printf("Other Count: %d\n", otherCnt)
}