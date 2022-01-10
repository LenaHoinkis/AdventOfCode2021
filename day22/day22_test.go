package main

import (
	"testing"
)

func Test_volCuboid(t *testing.T) {
	A := *NewRange(10, 12, 10, 12, 10, 12)
	B := *NewRange(12, 12, 12, 12, 12, 12)
	cubes := [...]cubeRange{A, B}
	vol := volCuboids(cubes[:])
	if vol != 28 {
		t.Errorf("Expected 28 got: %d", vol)
	}
}

func Test_Part2(t *testing.T) {
	A := *NewRange(10, 12, 10, 12, 10, 12)
	B := *NewRange(11, 13, 11, 13, 11, 13)
	C := *NewRange(10, 10, 10, 10, 10, 10)
	cubes := []cubeRange{A, B, C}
	if solvePart2(cubes) != 46 {
		t.Errorf("expected 46 got %d", solvePart2(cubes))
	}
}
