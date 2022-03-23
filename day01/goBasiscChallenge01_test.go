package main

import (
	"math"
	"testing"
)

func TestCalcCircleArea(t *testing.T) {
	result := calcCircleArea(1)
	if result != math.Pi {
		t.Errorf("Incorrect result: %f", result)
	}
}
