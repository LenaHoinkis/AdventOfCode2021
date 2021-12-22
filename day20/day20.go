package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/lenahoinkis/AdventOfCode2021/utils"
)

var inputFile = flag.String("inputFile", "data.input", "Relative file path to use as input.")

//Todo this onlz works with my data import where its changeing constantly
var dirtyGlobalVariable = '.'

func main() {
	flag.Parse()
	lines, _ := utils.ReadLines(*inputFile)
	code := lines[0]
	inputImage := lines[2:]

	height := len(inputImage)
	width := len(inputImage[0])

	//init
	image := make([]rune, width*height)
	for x, line := range inputImage {
		for y, c := range line {
			image[(x)*width+(y)] = c
		}
	}

	for i := 0; i < 50; i++ {
		width += 4
		height += 4
		image = imageIteration(image, width, height, code)
		width -= 2
		height -= 2
	}
	fmt.Println(sumImage(image))
}

func imageIteration(image []rune, width int, height int, code string) []rune {
	nextImage := createPadding(image, width, height)
	copyNextImage := make([]rune, len(nextImage))
	copy(copyNextImage, nextImage)

	for x := 0; x < width-2; x++ {
		for y := 0; y < height-2; y++ {
			i := x*width + 1*width + y + 1
			fieldValue := getFieldValue(nextImage, i, width)
			copyNextImage[i] = rune(code[fieldValue])
		}
	}
	return removePadding(copyNextImage, width, height)
}

func getFieldValue(image []rune, i int, row int) int {
	var sb strings.Builder
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			newIndex := ((x)*row + (y)) + i
			if image[newIndex] == '#' {
				sb.WriteString("1")
			} else {
				sb.WriteString("0")
			}
		}
	}
	result, _ := utils.BinaryToInt(sb.String())
	return int(result)
}

func createPadding(image []rune, width int, height int) []rune {
	newImage := make([]rune, width*width)
	//TODO: if 0 = # we need to chane it after the first iteration
	for i := range newImage {
		newImage[i] = dirtyGlobalVariable
	}

	if dirtyGlobalVariable == '.' {
		dirtyGlobalVariable = '#'
	} else {
		dirtyGlobalVariable = '.'
	}

	//put it in the mid
	for x := 0; x < width-4; x++ {
		for y := 0; y < height-4; y++ {
			newImage[(x+2)*width+(y+2)] = image[x*(width-4)+y]
		}
	}
	return newImage
}

func removePadding(image []rune, width int, height int) []rune {
	newImage := make([]rune, (width-2)*(width-2))
	//put it in the mid
	for x := 0; x < width-2; x++ {
		for y := 0; y < height-2; y++ {
			newImage[(x)*(width-2)+(y)] = image[(x+1)*(width)+y+1]
		}
	}
	return newImage
}

func printImage(image []rune, width int) {
	for i, v := range image {
		fmt.Print(string(v))
		if (i+1)%width == 0 {
			fmt.Println()
		}
	}
}

func sumImage(image []rune) int {
	sum := 0
	for _, v := range image {
		if v == '#' {
			sum++
		}
	}
	return sum
}
