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
	on     bool
}

func NewRange(x1, x2, y1, y2, z1, z2 int) *cubeRange {
	cr := &cubeRange{x1, x2, y1, y2, z1, z2, true}
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

//type cuberanges []*cubeRange

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

var inputFile = flag.String("inputFile", "data.input", "Relative file path to use as input.")

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

	//thanks to the solution of https://github.com/bozdoz/advent-of-code-2021/blob/main/22/cubes.go
	//I used the same approach but failed :(
	fmt.Println("Part2")
	var cuboids []cubeRange
	for _, line := range lines {
		isOn, cuboid := readInputLine(line)
		cuboid.on = isOn
		cuboids = append(cuboids, cuboid)
	}
	fmt.Println(solvePart2(cuboids))
}

func solvePart2(cuboids []cubeRange) int {
	result := 0

	for i := len(cuboids) - 1; i >= 0; i-- {
		cube := (cuboids)[i]

		if !cube.on {
			continue
		}

		intersections := []cubeRange{}

		// get all overlapping cubes (forwards)
		for _, next := range (cuboids)[i+1:] {
			intersection := getIntersection(next, cube)

			if intersection == nil {
				// did not intersect
				continue
			}

			// in recursive calls, "isOn" is synonymous with "shouldCountVolume"
			// as in, even if it's "off" we should calculate the total volume
			// of the intersections
			shouldCountVolume := true
			intersection.on = shouldCountVolume

			// if there is an intersection save it, and
			// reverse all intersections of the cubes intersections
			// i.e. don't count intersecting parts twice, and don't
			// discount intersecting parts twice.
			intersections = append(intersections, *intersection)
		}

		result += volCuboid(cube)
		result -= solvePart2(intersections)

	}
	return result
}

func volCuboids(cs []cubeRange) int {
	vol := 0
	for _, c := range cs {
		vol += volCuboid(c)
	}
	return vol
}

func volCuboid(cr cubeRange) int {
	return (cr.endX - cr.startX + 1) * (cr.endY - cr.startY + 1) * (cr.endZ - cr.startZ + 1)
}

func getIntersection(a cubeRange, b cubeRange) *cubeRange {
	x1 := max(a.startX, b.startX)
	x2 := min(a.endX, b.endX)
	y1 := max(a.startY, b.startY)
	y2 := min(a.endY, b.endY)
	z1 := max(a.startZ, b.startZ)
	z2 := min(a.endZ, b.endZ)

	if (x2 < x1) || y2 < y1 || z2 < z1 {
		return nil
	}

	return &cubeRange{x1, x2, y1, y2, z1, z2, a.on}
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
