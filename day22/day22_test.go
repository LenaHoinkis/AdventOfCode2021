package main

import (
	"fmt"
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

func Test_getAllIntersections(t *testing.T) {
	A := *NewRange(10, 12, 10, 12, 10, 12)
	B := *NewRange(12, 12, 12, 12, 12, 12)
	C := *NewRange(11, 12, 11, 12, 11, 12)
	cubes := [...]cubeRange{A, B, C}
	intersections := getAllIntersections(cubes[:], 2)
	if len(intersections) != 3 {
		t.Errorf("Expected 3 intersections AuB AuC BuC got: %d", len(intersections))
	}
	vol := volCuboids(intersections[:])
	if vol != 1+1+8 {
		t.Errorf("Expected 10 got: %d", vol)
	}
}

func Test_getAllIntersections_multiple(t *testing.T) {
	A := *NewRange(10, 12, 10, 12, 10, 12)
	B := *NewRange(12, 12, 12, 12, 12, 12)
	C := *NewRange(11, 12, 11, 12, 11, 12)
	cubes := [...]cubeRange{A, B, C}
	intersections := getAllIntersections(cubes[:], 2)
	if len(intersections) != 3 {
		t.Errorf("Expected 3 intersections AuB AuC BuC got: %d", len(intersections))
	}
	vol := volCuboids(intersections[:])
	if vol != 1+1+8 {
		t.Errorf("Expected 10 got: %d", vol)
	}
	newIntersection := getAllIntersections(intersections, 3)
	if len(newIntersection) != 1 {
		t.Errorf("Expected 1 intersections AuBuC got: %d", len(newIntersection))
	}
	vol = volCuboids(newIntersection[:])
	if vol != 1 {
		t.Errorf("Expected 1 got: %d", vol)
	}
}

func Test_getAllIntersections_4(t *testing.T) {
	A := *NewRange(10, 12, 10, 12, 10, 12)
	B := *NewRange(16, 14, 16, 14, 16, 14)
	C := *NewRange(11, 12, 11, 12, 11, 12)
	D := *NewRange(12, 12, 12, 12, 12, 12)
	cubes := [...]cubeRange{A, B, C, D}
	intersections := getAllIntersections(cubes[:], 4)
	if len(intersections) != 1 {
		t.Errorf("Expected 1 intersections AuBuC got: %d", len(intersections))
	}
	vol := volCuboids(intersections[:])
	if vol != 1 {
		t.Errorf("Expected 1 got: %d", vol)
	}
}

func Test_Pool(t *testing.T) {
	tuple := 3
	A := *NewRange(10, 12, 10, 12, 10, 12)
	B := *NewRange(12, 12, 12, 12, 12, 12)
	C := *NewRange(11, 12, 11, 12, 11, 12)
	D := *NewRange(10, 12, 10, 12, 10, 12)
	cubes := []cubeRange{A, B, C, D}
	p := Pool(tuple, cubes)
	if len(p) != 4 {
		t.Errorf("Expected 4 intersections ABC ABD BCD ACD got: %d", len(p))
	}
}

func Test_Deletion(t *testing.T) {
	A := *NewRange(10, 12, 10, 12, 10, 12)
	B := *NewRange(11, 13, 11, 13, 11, 13)
	L := *NewRange(11, 12, 11, 12, 11, 12)
	cubes := []cubeRange{A, B, L}
	//L,A,B (nothing happens)

	//A,L,B (nothing happens B fills the gap)

	//B,L,A (L deletes parts of B)

	//A,B,L (L deletes parts of B and the intersection between AB)
	fmt.Println(cubes)
}
