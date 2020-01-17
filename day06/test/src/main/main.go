package main

import "fmt"

type Car interface {
	GetName() string
	Run()
	DiDi()
}

type BMW struct {
	Name string
}

func (bmw *BMW) GetName() string {
	return bmw.Name
}

func (bmw *BMW) Run() {
	fmt.Printf("%s is run.\n", bmw.GetName())
}

func (bmw *BMW) DiDi() {
	fmt.Printf("DiDi from %s.\n", bmw.GetName())
}

func main() {
	var car Car
	var bmw BMW
	bmw.Name = "BMW"
	car = &bmw

	car.Run()
}
