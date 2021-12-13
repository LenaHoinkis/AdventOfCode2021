package main

import (
	"flag"
	"fmt"

	"github.com/lenahoinkis/AdventOfCode2021/utils"
)

func main() {
	inputFile := flag.String("inputFile", "t.input", "Relative file path to use as input.")
	flag.Parse()
	lines, _ := utils.ReadIntsMatrix(*inputFile)
	fmt.Println(lines)

	//part1
	flashes := 0
	flashesSum := 0
	for i := 0; i < 100; i++ {
		lines = increaseStep(lines)
		lines, flashes = flashStep(lines)
		flashesSum += flashes
	}
	fmt.Println(flashesSum)

	//part2
	step := 0
	for !checkAllZeros(lines) {
		step++
		lines = increaseStep(lines)
		lines, _ = flashStep(lines)
	}
	fmt.Println(step + 100) //dirty, but we have to add the 100 steps from part1

}
func checkAllZeros(oct [][]int) bool {
	for _, line := range oct {
		for _, o := range line {
			if o != 0 {
				return false
			}
		}
	}
	return true
}

func increaseStep(oct [][]int) [][]int {
	for x, line := range oct {
		for y, o := range line {
			oct[x][y] = o + 1
		}
	}
	return oct
}

func flashStep(oct [][]int) ([][]int, int) {
	flashes := 0
	flashesSum := 0
	for x, line := range oct {
		for y := range line {
			oct, flashes = flashIfPossible(x, y, oct, 0)
			flashesSum += flashes
		}
	}
	return oct, flashesSum
}

// increase neighbors and flash
func flashIfPossible(x, y int, oct [][]int, flashes int) ([][]int, int) {
	if oct[x][y] == 10 {
		oct[x][y] = 11 //we already worked on that
		//this loops is the same then copy all 9 neighbours
		for x1 := x - 1; x1 < x+2; x1++ {
			for y1 := y - 1; y1 < y+2; y1++ {
				if !(x == x1 && y == y1) {
					oct, flashes = flashNeighbours(x1, y1, oct, flashes)
				}
			}
		}
		oct[x][y] = 0
		return oct, flashes + 1
	}
	return oct, flashes
}
func flashNeighbours(x, y int, oct [][]int, flashes int) ([][]int, int) {
	if x < 0 || y < 0 || x >= len(oct) || y >= len(oct[x]) || oct[x][y] == 11 || oct[x][y] == 0 {
		return oct, flashes
	}
	if oct[x][y] != 10 {
		oct[x][y]++
	}
	return flashIfPossible(x, y, oct, flashes)
}
