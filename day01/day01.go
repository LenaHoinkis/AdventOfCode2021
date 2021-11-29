package main

import (
	"flag"
	"fmt"

	"github.com/lenahoinkis/AdventOfCode2021/utils"
)

var inputFile = flag.String("inputFile", "day01.input", "Relative file path to use as input.")

func main() {
	int, err := utils.ReadLinesOfInt(*inputFile)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(int)

}
