package main

import (
	"fmt"
	"strings"
)

func pathProcess(url string) string {
	if !strings.HasSuffix(url, "/") {
		url = fmt.Sprintf("%s/", url)
	}
	return url
}

func main() {
	var str string
	fmt.Scanf("%s", &str)
	str = pathProcess(str)
	fmt.Println(str)
}