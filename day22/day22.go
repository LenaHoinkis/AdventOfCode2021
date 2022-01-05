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

var inputFile = flag.String("inputFile", "ex3.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	lines, _ := utils.ReadLines(*inputFile)

	//part1
	max := 50
	min := -50
	c := NewCube(max, min)
	for _, line := range lines {
		setOn, x1, x2, y1, y2, z1, z2 := readInputLine(line)
		if !checkOutOfBound(x1, x2, y1, y2, z1, z2) {
			setField(setOn, c, cubeRange{x1, x2, y1, y2, z1, z2})
		}
	}
	fmt.Println(countOn(c))

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

func readInputLine(line string) (bool, int, int, int, int, int, int) {
	var x1, x2, y1, y2, z1, z2 int
	setOn := true
	if line[2] == 'f' {
		fmt.Sscanf(line, "off x=%d..%d,y=%d..%d,z=%d..%d", &x1, &x2, &y1, &y2, &z1, &z2)
		setOn = false
	} else {
		fmt.Sscanf(line, "on x=%d..%d,y=%d..%d,z=%d..%d", &x1, &x2, &y1, &y2, &z1, &z2)
	}
	return setOn, x1, x2, y1, y2, z1, z2
}
