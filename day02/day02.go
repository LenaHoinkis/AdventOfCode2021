package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/lenahoinkis/AdventOfCode2021/utils"
)

var inputFile = flag.String("inputFile", "day02.input", "Relative file path to use as input.")

func main() {
	lines, err := utils.ReadLines(*inputFile)
	if err != nil {
		fmt.Print(err)
	}

	//part1
	x, y := 0, 0
	for _, v := range lines {
		str := strings.Split(v, " ")
		i, err := strconv.Atoi(string(str[1]))
		if err != nil {
			panic(1)
		}
		switch str[0] {
		case "forward":
			x += i
		case "down":
			y += i
		case "up":
			y -= i
		}
	}
	fmt.Println(y * x)

	//part2
	x, y = 0, 0
	aim := 0
	for _, v := range lines {
		str := strings.Split(v, " ")
		i, err := strconv.Atoi(string(str[1]))
		if err != nil {
			panic(1)
		}
		switch str[0] {
		case "forward":
			x += i
			y += i * aim
		case "down":
			aim += i
		case "up":
			aim -= i
		}
	}
	fmt.Println(y * x)
}
