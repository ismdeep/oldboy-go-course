package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "hello")
}

func main() {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		fmt.Println("Http Listen Failed. err:", err)
	}
}
