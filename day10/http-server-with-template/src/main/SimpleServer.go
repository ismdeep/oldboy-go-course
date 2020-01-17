package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("Parse file err:", err)
		return
	}
	p := Person{
		Name: "ismdeep",
		Age:  "18",
	}
	if err := t.Execute(w, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
