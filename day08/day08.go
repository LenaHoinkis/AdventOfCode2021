package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Note struct {
	signals []string
	outputs []string
}

func main() {
	inputFile := flag.String("inputFile", "ex.input", "Relative file path to use as input.")
	flag.Parse()
	notes, _ := ReadLinesNotes(*inputFile)

	//part1
	count1478 := 0
	for _, note := range notes {
		for _, word := range note.outputs {
			// 1->2 4->4 7->3 8->7
			numberOfChars := len(word)
			if numberOfChars == 2 || numberOfChars == 4 || numberOfChars == 3 || numberOfChars == 7 {
				count1478++
			}
		}
	}
	fmt.Println(count1478)

	//part2
	//fill a binary map
	//m := make(map[int][7]int)
	sum := 0
	for _, note := range notes {
		m := make(map[int][7]int)

		for _, word := range note.signals {
			numberOfChars := len(word)
			for _, char := range word {
				slice := m[numberOfChars]
				slice[char-97]++
				m[numberOfChars] = slice
			}
		}

		// I calculated the required steps on paper
		// by subtracting our constants 1->2 4->4 7->3 8->7 step by step
		var code [7]int
		code[0] = findPos(1, calcDifference(m[3], m[2], 0))
		deleteFromMapSlice(code[0], m)
		code[4] = findPos(1, calcDifference(m[5], m[4], 2))
		deleteFromMapSlice(code[4], m)
		code[6] = findPos(3, calcDifference(m[6], m[4], 0))
		deleteFromMapSlice(code[6], m)
		code[3] = findPos(3, calcDifference(m[5], m[3], 0))
		deleteFromMapSlice(code[3], m)
		code[1] = findPos(1, calcDifference(m[5], m[3], 2))
		deleteFromMapSlice(code[1], m)
		code[5] = findPos(1, calcDifference(m[6], m[2], 1))
		deleteFromMapSlice(code[5], m)
		code[2] = findPos(2, m[5])

		decoded := ""
		for _, word := range note.outputs {
			var slice [7]int
			for _, char := range word {
				slice[findPos(int(char-97), code)]++
			}
			switch slice {
			case [7]int{1, 1, 1, 0, 1, 1, 1}:
				decoded += "0"
			case [7]int{0, 0, 1, 0, 0, 1, 0}:
				decoded += "1"
			case [7]int{1, 0, 1, 1, 1, 0, 1}:
				decoded += "2"
			case [7]int{1, 0, 1, 1, 0, 1, 1}:
				decoded += "3"
			case [7]int{0, 1, 1, 1, 0, 1, 0}:
				decoded += "4"
			case [7]int{1, 1, 0, 1, 0, 1, 1}:
				decoded += "5"
			case [7]int{1, 1, 0, 1, 1, 1, 1}:
				decoded += "6"
			case [7]int{1, 0, 1, 0, 0, 1, 0}:
				decoded += "7"
			case [7]int{1, 1, 1, 1, 1, 1, 1}:
				decoded += "8"
			case [7]int{1, 1, 1, 1, 0, 1, 1}:
				decoded += "9"
			}
		}
		x, _ := strconv.Atoi(decoded)
		sum += x
	}
	fmt.Println(sum)
}

func calcDifference(x [7]int, y [7]int, factor int) [7]int {
	r := x
	for i := 0; i <= factor; i++ {
		for i := range x {
			r[i] = r[i] - y[i]
		}
	}
	return r
}

func deleteFromMapSlice(pos int, m map[int][7]int) {
	for i, slice := range m {
		slice[pos] = 0
		m[i] = slice
	}
}

func findPos(n int, x [7]int) int {
	for i, v := range x {
		if v == n {
			return i
		}
	}
	return 0
}

//ReadLinesOfInt reads line and convert the number to int
func ReadLinesNotes(path string) ([]Note, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var notes []Note
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var note Note
		line := scanner.Text()
		singalsAndOutputs := strings.Split(line, " | ")
		note.signals = append(note.signals, strings.Split(singalsAndOutputs[0], " ")...)
		note.outputs = append(note.outputs, strings.Split(singalsAndOutputs[1], " ")...)
		notes = append(notes, note)
	}
	return notes, scanner.Err()
}
