package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", SimpleServer)
	fmt.Println("Starting server...")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("Http Listen Failed, err:", err)
	}

}
