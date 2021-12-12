package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/lenahoinkis/AdventOfCode2021/utils"
	"github.com/pkg/errors"
)

func main() {
	inputFile := flag.String("inputFile", "ex.input", "Relative file path to use as input.")
	flag.Parse()
	lines, _ := utils.ReadLines(*inputFile)

	//part1
	pointsSum := 0
	//part2
	var scores []int

	for _, line := range lines {
		var err error
		for i := 1; i < len(line); i++ {
			//part1 find all corrupt lines
			var points int
			switch string(line[i]) {
			case "]":
				i, line, err = checkSyntax("[", line, i)
				points = 57
			case "}":
				i, line, err = checkSyntax("{", line, i)
				points = 1197
			case ")":
				i, line, err = checkSyntax("(", line, i)
				points = 3
			case ">":
				i, line, err = checkSyntax("<", line, i)
				points = 25137
			}
			if err != nil {
				fmt.Println(err)
				pointsSum += points
				break
			}
		}
		//part2 for all uncomplete lines
		if err == nil {
			pointsPart2 := 0
			for i := len(line) - 1; i >= 0; i-- {
				pointsPart2 *= 5
				switch string(line[i]) {
				case "[":
					pointsPart2 += 2
				case "{":
					pointsPart2 += 3
				case "(":
					pointsPart2 += 1
				case "<":
					pointsPart2 += 4
				}
			}
			scores = append(scores, pointsPart2)
		}
	}
	fmt.Println(pointsSum)
	sort.Ints(scores)
	mid := int(len(scores) / 2)
	fmt.Println(scores[mid])
}

func checkSyntax(Open string, line string, i int) (int, string, error) {
	if string(line[i-1]) == Open {
		line = line[0:i-1] + line[i+1:]
		i = i - 2
	} else {
		return i, line, errors.Errorf("Error at %d /n", i)
	}
	return i, line, nil
}
