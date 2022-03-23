package main

import "fmt"

type Astro struct {
	name        string
	age         int
	lastMission string
	isNeeded    bool
}

type NasaMission struct {
	people  []Astro
	number  int
	message string
}

func main() {
	astro1 := Astro{age: 1, name: "One", lastMission: "mis1", isNeeded: true}
	astro2 := Astro{age: 2, name: "two", lastMission: "mis2", isNeeded: true}
	astro3 := Astro{age: 3, name: "three", lastMission: "SpaceX Crew-3", isNeeded: true}

	fmt.Println(astro1)
	fmt.Println(astro2)
	fmt.Println(astro3)

	slice := []Astro{astro1, astro2, astro3}
	fmt.Println("------------")
	fmt.Println(slice)
	fmt.Println("------------")
	fmt.Println(slice[2])

	fmt.Println("############")
	s := NasaMission{people: slice, number: 3, message: "success"}
	// display "s" without fields
	fmt.Println(s)

	// display "s" with fields
	fmt.Printf("%+v", s)

	fmt.Println("############")
	s2 := NasaMission{number: 3, message: "fail"}
	fmt.Println(s2)

}
