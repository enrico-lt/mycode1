package main

import "fmt"

type Car struct {
	color string
}

func NewCar(color string) *Car {
	car := Car{color: color}
	return &car
}

func main() {

	c := NewCar("red")
	fmt.Println("Car1: ", c)

}
