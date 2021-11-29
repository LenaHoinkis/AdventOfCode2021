package utils

import (
	"bufio"
	"os"
	"strconv"
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
