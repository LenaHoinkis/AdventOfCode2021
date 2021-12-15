package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//ReadLinesOfInt reads line and convert the number to int
func ReadLinesOfInt(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}
	return numbers, scanner.Err()
}

//ReadLinesOfInt reads line and convert the number to int
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var result []int
	intString := scanner.Text()
	for _, v := range strings.Split(intString, ",") {
		x, err := strconv.Atoi(v)
		if err != nil {
			return result, err
		}
		result = append(result, x)

	}
	return result, scanner.Err()
}

func ReadIntsMatrix(path string) ([][]int, error) {
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
