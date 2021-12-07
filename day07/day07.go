package main

import (
	"flag"
	"fmt"
	"math"
	"sort"

	"github.com/lenahoinkis/AdventOfCode2021/utils"
)

func main() {
	inputFile := flag.String("inputFile", "ex.input", "Relative file path to use as input.")
	flag.Parse()

	crabs, _ := utils.ReadInts(*inputFile)

	sort.Ints(crabs)

	//part1
	//with for loop and !gettingBigger as additional condition
	smallestDistance := 0
	gettingBigger := false
	for i := 0; !gettingBigger && i < crabs[len(crabs)-1]; i++ {
		distance := 0
		for x := 0; !gettingBigger && x < len(crabs); x++ {
			distance += int(math.Abs(float64(crabs[x] - i)))
			if i > 0 && distance > smallestDistance {
				gettingBigger = true
			}
		}
		if !gettingBigger {
			smallestDistance = distance
		}
	}
	fmt.Println(smallestDistance)

	//part2
	//with range and more if conditions to break
	smallestDistance = 0
	for i := 0; i < crabs[len(crabs)-1]; i++ {
		distanceGaus := 0
		gettingBigger := false
		for _, pos := range crabs {
			distance := math.Abs(float64(pos - i))
			//The Gaus calculation is the difference of part2
			distanceGaus += int((math.Pow(distance, 2) + float64(distance)) / 2)
			if i > 0 && distanceGaus > smallestDistance {
				gettingBigger = true
				break
			}
		}
		if gettingBigger {
			break
		} else {
			smallestDistance = distanceGaus
		}
	}
	fmt.Println(smallestDistance)
}
