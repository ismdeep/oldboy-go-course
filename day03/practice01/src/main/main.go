package main

import (
	"fmt"
	"strings"
)

func urlProcess(url string) string {
	if !strings.HasPrefix(url, "https://") {
		url = fmt.Sprintf("https://%s", url)
	}
	return url
}

func main() {
	var str string
	fmt.Scanf("%s", &str)
	str = urlProcess(str)
	fmt.Println(str)
}
