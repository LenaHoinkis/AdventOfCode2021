package main

import (
	"flag"
	"fmt"

	"github.com/lenahoinkis/AdventOfCode2021/utils"
)

func main() {
	inputFile := flag.String("inputFile", "ex.input", "Relative file path to use as input.")
	flag.Parse()

	fishs, _ := utils.ReadInts(*inputFile)

	//part1
	for i := 0; i < 80; i++ {
		for i, fish := range fishs {
			if fish == 0 {
				fish = 6
				fishs = append(fishs, 8)
			} else {
				fish--
			}
			fishs[i] = fish
		}
	}
	fmt.Println(len(fishs))

	//part2
	initalFishs, _ := utils.ReadInts(*inputFile)
	fish := make([]int, 9)
	for _, v := range initalFishs { //inital fill
		fish[v]++
	}
	for i := 0; i < 256; i++ {
		fishToBreed := fish[0]           // Remember fish at 0 days
		fish = fish[1:]                  // 'Pop' fish at 0 days
		fish[6] += fishToBreed           // Add fish to 6th day
		fish = append(fish, fishToBreed) // Add fish to 8th day
	}
	fmt.Printf("Part2: %d\n", getTotalFish(fish))
}

func getTotalFish(fish []int) int {
	total := 0
	for _, val := range fish {
		total += val
	}
	return total
}
