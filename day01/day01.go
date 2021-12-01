package main

import (
	"flag"
	"fmt"

	"github.com/lenahoinkis/AdventOfCode2021/utils"
)

var inputFile = flag.String("inputFile", "day01.input", "Relative file path to use as input.")

func main() {
	depths, err := utils.ReadLinesOfInt(*inputFile)
	if err != nil {
		fmt.Print(err)
	}
	// Part 1
	count := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			count++
		}
	}
	fmt.Println(count)

	//Part 2
	count = 0
	var sumOld int
	for i := 2; i < len(depths); i++ {
		sumNew := depths[i] + depths[i-1] + depths[i-2]
		if i > 2 && sumNew > sumOld {
			count++
		}
		sumOld = sumNew
	}
	fmt.Println(count)
}
