package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
type Rule struct {
	Pair   string
	Insert string
}*/

func main() {
	inputFile := flag.String("inputFile", "data.input", "Relative file path to use as input.")
	flag.Parse()
	template1, m, _ := pattern(*inputFile)
	template2 := template1

	//with stringBuilder its fast enough for 26 loops
	//with reassigning template each step it only worked for 18
	//idea would be to use somehow regex
	//or to add a new longer pattern to the map
	//or both
	var sb strings.Builder
	for i := 0; i < 5; i++ {
		for i := 0; i < len(template1)-1; i++ {
			sb.WriteString(string(template1[i]))
			add, ok := m[template1[i:i+2]]
			if ok {
				sb.WriteString(add)
			}
		}
		sb.WriteString(string(template1[len(template1)-1]))
		template1 = sb.String()
		sb.Reset()
	}
	charmap := make(map[rune]int)
	var min, max int
	for _, v := range template1 {
		charmap[v]++
		min, max = charmap[v], charmap[v]
	}
	for _, v := range charmap {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	fmt.Println(max - min)

	//See Lanternfish
	fmt.Println("part2")

	//pairs count all our pair combination
	//we can later lookup in m what character gets inserted
	pairs := make(map[string]int)

	//we save the characters only to have the result in the end
	//it would also be possible to calc them by the pairs map we saved
	characters := make(map[string]int)

	//inital fill
	//counting all pairs (with overlap) and all characters
	for i := 0; i < len(template2)-1; i++ {
		pair := template2[i : i+2]
		pairs[pair]++
		characters[template2[i:i+1]]++
	}
	characters[template2[len(template2)-1:]]++

	for i := 0; i < 40; i++ {
		//we create a temp map to not manipulate our calculation before we went trough all combinations
		tempPairs := make(map[string]int)
		for pair, count := range pairs {
			//get the new char
			translate := m[pair]
			//when we insert a char in between we get two new combinations BN -> V will get BV and VN
			tempPairs[translate+pair[1:]] += count
			tempPairs[pair[:1]+translate] += count
			//save the new character for the result in the end
			characters[translate] += count
		}
		pairs = tempPairs
	}
	fmt.Println(maxMinusMin(characters))
}

func maxMinusMin(m map[string]int) int {
	max, min := m["B"], m["B"]
	for _, v := range m {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}

//Readfunction
func pattern(path string) (string, map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	var template string
	//var rules []Rule
	m := make(map[string]string)
	row := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if row == 0 {
			template = line
		}
		if row > 1 {
			s := strings.Split(line, " -> ")
			//rules = append(rules, Rule{s[0], s[1]})
			m[s[0]] = s[1]
		}
		row++
	}
	return template, m, scanner.Err()
}
