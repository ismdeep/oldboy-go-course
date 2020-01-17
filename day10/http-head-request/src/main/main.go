package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var urls = []string{
	"https://www.baidu.com",
	"https://www.google.cn",
	"https://taobao.com",
	"http://jw.jxust.edu.cn",
}

func main() {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			client := http.Client{
				Transport:     nil,
				CheckRedirect: nil,
				Jar:           nil,
				Timeout:       2 * time.Second,
			}
			resp, err := client.Head(url)
			if err != nil {
				fmt.Printf("Head %s failed, err: %v\n", url, err)
			} else {
				fmt.Printf("Head successed {%s}, Status: %v, Length: %d\n", url, resp.Status, resp.ContentLength)
			}
			wg.Add(-1)
		}(url)
	}
	wg.Wait()
}
