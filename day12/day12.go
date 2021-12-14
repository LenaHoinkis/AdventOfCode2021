package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	inputFile := flag.String("inputFile", "t.input", "Relative file path to use as input.")
	flag.Parse()
	m, _ := ReadPaths(*inputFile)
	m2 := copyMap(m)

	//we search the start
	fmt.Println(findPath("start", m))

	fmt.Println(findPathPart2("start", "", m2, false, make(map[string]int)))
}

func findPath(currentPoint string, m map[string][]string) int {
	//we are at the end
	if currentPoint == "end" {
		return 1
	}
	//copy to not delete values in recursion
	m1 := copyMap(m)

	//never back to small points
	if unicode.IsLower(rune(currentPoint[0])) {
		removeFromMap(currentPoint, m1)
	}
	//now go deep
	x := 0
	connections := m[currentPoint]
	for _, connection := range connections {
		x += findPath(connection, m1)
	}
	return x
}

func findPathPart2(currentPoint, currentPath string, m map[string][]string, usedSecondVisit bool, lookup map[string]int) int {
	currentPath = currentPath + currentPoint
	//we are at the end
	if currentPoint == "end" {
		_, ok := lookup[currentPath]
		if !ok {
			lookup[currentPath]++
			return 1
		}
		return 0
	}
	//copy to not delete values in recursion
	m1 := copyMap(m)

	x := 0
	connections := m[currentPoint]
	//go back to old places once or take a small cave twice
	if unicode.IsLower(rune(currentPoint[0])) {
		if !usedSecondVisit && currentPoint != "start" && currentPoint != "end" {
			for _, connection := range connections {
				x += findPathPart2(connection, currentPath, m1, true, lookup)
			}
		}
		removeFromMap(currentPoint, m1)
	}

	//now go deep
	for _, connection := range connections {
		x += findPathPart2(connection, currentPath, m1, usedSecondVisit, lookup)
	}
	return x
}

//I hate maps, thanks for wasting my time
func copyMap(m map[string][]string) map[string][]string {
	m1 := make(map[string][]string)
	for v, cs := range m {
		cs1 := make([]string, len(cs))
		copy(cs1, cs)
		m1[v] = cs1
	}
	return m1
}

func removeFromMap(p string, m map[string][]string) {
	for v, cs := range m {
		for i, c := range cs {
			if c == p {
				cs = remove(cs, i)
				break
			}
		}
		m[v] = cs
	}
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func ReadPaths(path string) (map[string][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	m := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, "-")
		m = addValueToMapSlice(points[0], points[1], m)
		m = addValueToMapSlice(points[1], points[0], m)
	}
	return m, scanner.Err()
}

func addValueToMapSlice(key, value string, m map[string][]string) map[string][]string {
	connections := m[key]
	connections = append(connections, value)
	m[key] = connections
	return m
}
