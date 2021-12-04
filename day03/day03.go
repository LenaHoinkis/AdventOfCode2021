package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/lenahoinkis/AdventOfCode2021/utils"
)

var inputFile = flag.String("inputFile", "ex.input", "Relative file path to use as input.")

func main() {
	lines, err := utils.ReadLines(*inputFile)
	if err != nil {
		fmt.Print(err)
	}

	lineLenght := len(lines[0])
	inputLenght := len(lines)
	var sums = make([]int, lineLenght)

	//part1
	for _, line := range lines {
		for i, char := range line {
			sums[i] += int(char) - 48
		}
	}
	gamma := ""
	epsilon := ""
	for _, sum := range sums {
		if sum >= inputLenght/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(g * e)

	//part2
	var oxygen, scrubber int64
	oxygen = calcValue(lines, false)
	scrubber = calcValue(lines, true)
	fmt.Println(oxygen * scrubber)
}

func calcValue(lines []string, takeZeros bool) int64 {
	var x int64
	for i := 0; i < len(lines[0])+1; i++ {
		if len(lines) == 1 {
			fmt.Println(lines)
			x, _ = strconv.ParseInt(lines[0], 2, 64)
			break
		}
		var zeros, ones []string
		for y := 0; y < len(lines); y++ {
			if string(lines[y][i]) == "0" {
				zeros = append(zeros, lines[y])
			} else {
				ones = append(ones, lines[y])
			}
		}
		if takeZeros {
			if len(zeros) < len(ones) {
				lines = zeros
			} else if len(zeros) > len(ones) {
				lines = ones
			} else {
				lines = zeros
			}
		} else {
			if len(zeros) > len(ones) {
				lines = zeros
			} else if len(zeros) < len(ones) {
				lines = ones
			} else {
				lines = ones
			}
		}
	}
	return x
}
