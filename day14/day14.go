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
	inputFile := flag.String("inputFile", "ex.input", "Relative file path to use as input.")
	flag.Parse()
	template1, m, _ := pattern(*inputFile)
	//template2 := template1

	//with stringBuilder its fast enough for 26 loops
	//with reassigning template each step it only worked for 18
	//idea would be to use somehow regex
	//or to add a new longer pattern to the map
	//or both
	var sb strings.Builder
	for i := 0; i < 25; i++ {
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
	/*
		for k, v := range m {
			re := regexp.MustCompile(k)
			findings := re.FindAllIndex([]byte(template2), -1)
			for _, finding := range findings {
				template2 = join(template1[:finding[0]+1], v, template1[finding[1]:])
			}
		}
		fmt.Println(template2)
	*/
}

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