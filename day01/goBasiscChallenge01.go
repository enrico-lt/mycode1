package main

import (
	"fmt"
	"math"
)

var height, width int
var radius int

func calcArea(h, w int) int {
	return h * w
}

func calcCircleArea(radius int) float64 {
	return float64(radius*radius) * math.Pi
}

func main() {

	height = 1
	width = 2

	radius = 1

	resultArea := calcArea(height, width)

	resultCircle := calcCircleArea(radius)

	fmt.Println("The resultArea is", resultArea)
	fmt.Println("The resultCircle is", resultCircle)
}
