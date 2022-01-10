package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/lenahoinkis/AdventOfCode2021/utils"
)

type cubeRange struct {
	startX int
	endX   int
	startY int
	endY   int
	startZ int
	endZ   int
}

func NewRange(x1, x2, y1, y2, z1, z2 int) *cubeRange {
	cr := &cubeRange{x1, x2, y1, y2, z1, z2}
	if x2 < x1 {
		cr.startX = x2
		cr.endX = x1
	}
	if y2 < y1 {
		cr.startY = y2
		cr.endY = y1
	}
	if z2 < z1 {
		cr.startZ = z2
		cr.endZ = z1
	}
	return cr

}

type cube struct {
	size, min int
	fields    []bool
}

func NewCube(max, min int) *cube {
	c := new(cube)
	size := float64(max) + math.Abs(float64(min))
	c.fields = make([]bool, int(math.Pow(size+1, 3)))
	c.size = int(size)
	if min < 0 {
		c.min = int(math.Abs(float64(min)))
	} else {
		c.min = 0
	}
	return c
}

var inputFile = flag.String("inputFile", "ex2.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	lines, _ := utils.ReadLines(*inputFile)

	//part1
	fmt.Println("Part1")
	max := 50
	min := -50
	c := NewCube(max, min)
	for _, line := range lines {
		setOn, cr := readInputLine(line)
		if !checkOutOfBound(cr.startX, cr.endX, cr.startY, cr.endY, cr.startZ, cr.endZ) {
			setField(setOn, c, cr)
		}
		fmt.Println(countOn(c))
	}
	fmt.Println(countOn(c))

	/*
		part 2 solution with Inclusion–exclusion
		when two cupids intersect, we get a new cupid which we need to subtract

		example we get A then delete (L) and then add B. Everything 9x9 and intersection 1
		1: v = 0
		2: v+= A -> save A
		3: v+= (-LuA) -> save -LuA
		4: v+= B - AuB - (-BuLuA) -> save B

		Add A, Add B, Delete L
		1: v = 0 (0)
		2: v+= A -> save A (9)
		3: v+= B - AuB -> save B (17)
		4: v+= (-LuA) + (-LuB) (16 wrong!) because LuA=LuB, I need to merge those I want to subtract
	*/
	fmt.Println("Part2")
	result, v := 0, 0
	var cubes []cubeRange
	for _, line := range lines {
		setOn, cr := readInputLine(line)
		if setOn {
			v, cubes = getMergedIntersections(setOn, cr, cubes)
			// save cube to identify intersactions later (step2)
			v += volCupid(cr)
			cubes = append(cubes, cr)
			result += v
		} else {
			v, cubes = getMergedIntersections(setOn, cr, cubes)
			result += v
		}

		fmt.Println(result)
	}

}

func getMergedIntersections(setOn bool, cr cubeRange, cubes []cubeRange) (int, []cubeRange) {
	volumeToAdd := 0
	var intersections []cubeRange

	//go through all cuboids and see where it overlaps
	for _, c := range cubes {
		var intersection *cubeRange
		if setOn {
			intersection = getIntersection(c, cr)
		} else {
			intersection = getIntersection(cr, c)
		}
		if intersection != nil {
			intersections = append(intersections, *intersection)
		}
	}

	//go trough all intersections and smaller them if required
	for range intersections {
		intersections = smaller(cr, intersections)
	}

	//go trough all intersections and merge overlapping sections
	i := 0
	for range intersections {
		intersections, i = merge(intersections, i)
	}

	//go trough all intersections and sum up
	for _, intersection := range intersections {
		volumeToAdd -= volCupid(intersection)
		if !setOn {
			cubes = append(cubes, intersection)
		}
	}

	return volumeToAdd, cubes
}

func smaller(a cubeRange, list []cubeRange) []cubeRange {
	for i := 0; i < len(list); i++ {
		if doesContain(list[i], a) {
			list = remove(i, list)
		}
	}
	return list
}

func merge(list []cubeRange, index int) ([]cubeRange, int) {
	for i := index + 1; i < len(list); i++ {
		if doesContain(list[index], list[i]) {
			list = remove(i, list)
		} else if doesContain(list[i], list[index]) {
			list = remove(index, list)
			return list, index
		}
	}
	return list, index + 1
}

func remove(i int, a []cubeRange) []cubeRange {
	// Remove the element at index i from a.
	a[i] = a[len(a)-1] // Copy last element to index i.
	//a[len(a)-1] = nil  // Erase last element (write zero value).
	a = a[:len(a)-1] // Truncate slice.
	return a
}

func doesContain(a cubeRange, b cubeRange) bool {
	return (a.startX <= b.startX && a.endX >= b.endX) &&
		(a.startY <= b.startY && a.endY >= b.endY) &&
		(a.startZ <= b.startZ && a.endZ >= b.endZ)
}

func getIntersection(a cubeRange, b cubeRange) *cubeRange {
	if doesIntersect(a, b) {
		return &cubeRange{
			max(a.startX, b.startX),
			min(a.endX, b.endX),
			max(a.startY, b.startY),
			min(a.endY, b.endY),
			max(a.startZ, b.startZ),
			min(a.endZ, b.endZ),
		}
	}
	return nil
}

func doesIntersect(a cubeRange, b cubeRange) bool {
	return (a.startX <= b.endX && a.endX >= b.startX) &&
		(a.startY <= b.endY && a.endY >= b.startY) &&
		(a.startZ <= b.endZ && a.endZ >= b.startZ)
}

func volCupid(cr cubeRange) int {
	return (cr.endX - cr.startX + 1) * (cr.endY - cr.startY + 1) * (cr.endZ - cr.startZ + 1)
}

func setField(value bool, c *cube, cR cubeRange) {
	for x := cR.startX; x <= cR.endX; x++ {
		for y := cR.startY; y <= cR.endY; y++ {
			for z := cR.startZ; z <= cR.endZ; z++ {
				//here we add also the c.min, to get the values in a positve range
				pos := x + c.min + +c.size*(y+c.min) + c.size*c.size*(z+c.min)
				c.fields[pos] = value
			}
		}
	}
}

func countOn(c *cube) int {
	sum := 0
	for _, v := range c.fields {
		if v {
			sum++
		}
	}
	return sum
}

func checkOutOfBound(nums ...int) bool {
	for _, num := range nums {
		if num > 50 || num < -50 {
			return true
		}
	}
	return false
}

func readInputLine(line string) (bool, cubeRange) {
	var x1, x2, y1, y2, z1, z2 int
	setOn := true
	if line[2] == 'f' {
		fmt.Sscanf(line, "off x=%d..%d,y=%d..%d,z=%d..%d", &x1, &x2, &y1, &y2, &z1, &z2)
		setOn = false
	} else {
		fmt.Sscanf(line, "on x=%d..%d,y=%d..%d,z=%d..%d", &x1, &x2, &y1, &y2, &z1, &z2)
	}
	return setOn, *NewRange(x1, x2, y1, y2, z1, z2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
