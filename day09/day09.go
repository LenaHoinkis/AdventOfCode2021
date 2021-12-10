package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
)

func main() {
	inputFile := flag.String("inputFile", "ex.input", "Relative file path to use as input.")
	flag.Parse()
	hm, _ := ReadHeatmap(*inputFile)

	//give the heatmap a 9er padding
	height := len(hm) + 2
	width := len(hm[0]) + 2
	hmPadded := make([][]int, height)
	for y := range hmPadded {
		hmPadded[y] = make([]int, width)
		for x := range hmPadded[y] {
			hmPadded[y][x] = 9
		}
	}
	for y := 1; y <= len(hm); y++ {
		for x := 1; x <= len(hm[0]); x++ {
			hmPadded[y][x] = hm[y-1][x-1]
		}
	}

	//part1
	sumOfRiskFields := 0
	//part2
	var biggestFields [3]int
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if isRiksy(hmPadded, y, x) {
				//part1
				sumOfRiskFields += hmPadded[y][x] + 1
				//part2
				result := checkNeighbours(hmPadded, y, x, 0)
				if biggestFields[0] < result {
					biggestFields[0] = result
					sort.Ints(biggestFields[:])
				}
			}
		}
	}
	fmt.Println(sumOfRiskFields)
	fmt.Println(biggestFields[0] * biggestFields[1] * biggestFields[2])
}

func isRiksy(hmPadded [][]int, y int, x int) bool {
	return (hmPadded[y][x] < hmPadded[y][x+1] && hmPadded[y][x] < hmPadded[y+1][x] &&
		hmPadded[y][x] < hmPadded[y][x-1] && hmPadded[y][x] < hmPadded[y-1][x])
}
func checkNeighbours(hmPadded [][]int, y int, x int, sum int) int {
	if hmPadded[y][x] == 9 {
		return 0
	}

	hmPadded[y][x] = 9
	a, b, c, d := 0, 0, 0, 0
	a += checkNeighbours(hmPadded, y+1, x, sum)
	b += checkNeighbours(hmPadded, y-1, x, sum)
	c += checkNeighbours(hmPadded, y, x+1, sum)
	d += checkNeighbours(hmPadded, y, x-1, sum)
	return sum + 1 + a + b + c + d
}

func ReadHeatmap(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var intSlice []int
		for _, v := range line {
			intSlice = append(intSlice, int(v)-48)
		}
		result = append(result, intSlice)
	}
	return result, scanner.Err()
}
