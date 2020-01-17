package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str1 string
	str := "  ismdeep  \n "
	str1 = strings.TrimSpace(str)
	fmt.Printf("str = {%s}\n", str1)
	str1 = strings.Trim(str, " \np")
	fmt.Printf("str = {%s}\n", str1)
	str1 = strings.TrimLeft(str, " \np")
	fmt.Printf("str = {%s}\n", str1)
	str1 = strings.TrimRight(str, " \np")
	fmt.Printf("str = {%s}\n", str1)

	str = "httpsh://"
	fmt.Println(strings.TrimLeft(str, "https"))

	str = "  hello world.   "
	fmt.Println(strings.Fields(str))
	fmt.Printf("%v\n", strings.Fields(str))

	str = "a|b|c"
	fmt.Println(strings.Split(str, "|"))
	fmt.Println(strings.Join(strings.Split(str, "|"), "/"))

	var n int
	n = 203984
	str1 = strconv.Itoa(n)
	fmt.Println(str1)

	n, _ = strconv.Atoi("230480")
	fmt.Println(n)
}