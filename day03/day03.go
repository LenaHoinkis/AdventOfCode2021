package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/lenahoinkis/AdventOfCode2021/utils"
)

var inputFile = flag.String("inputFile", "day03ex.input", "Relative file path to use as input.")

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
}
