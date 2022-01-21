package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/lenahoinkis/AdventOfCode2021/utils"
)

func main() {
	inputFile := flag.String("inputFile", "ex1.input", "Relative file path to use as input.")
	flag.Parse()
	lines, _ := utils.ReadLines(*inputFile)
	input := lines[0]

	//hex2int
	s := utils.HexToBinary(input)
	fmt.Println(readPackage(s))
}

func readPackage(s string) (version int, newcontent string, r2 int) {
	v, _ := strconv.ParseInt(s[:3], 2, 64)
	version = int(v)
	id, _ := strconv.ParseInt(s[3:6], 2, 64)

	if id == 4 {
		newcontent, r2 = readLiteral(s[6:])
	} else {
		v := 0
		v, newcontent, r2 = readOperator(s[6:], int(id))
		version += v
	}
	return version, newcontent, r2
}

func readLiteral(content string) (string, int) {
	var sb strings.Builder
	i := 0
	for {
		sb.WriteString(content[i+1 : i+5])
		if content[i] == '0' {
			i = i + 5
			break
		}
		i = i + 5
	}
	v, _ := utils.BinaryToInt(sb.String())
	return content[i:], int(v)
}

func readOperator(content string, id int) (int, string, int) {
	versionPart2 := make([]int, 0)
	version := 0
	if content[0] == '0' {
		totalLenght, _ := utils.BinaryToInt(content[1:16]) // bits length
		newcontent := content[16 : totalLenght+16]
		for i := 0; len(newcontent) > 0; i++ {
			v, r2 := 0, 0
			v, newcontent, r2 = readPackage(newcontent)
			version += v
			versionPart2 = append(versionPart2, r2)
		}
		content = content[totalLenght+16:]
	} else {
		subPackageCount, _ := utils.BinaryToInt(content[1:12]) //11bits
		newcontent := content[12:]
		for i := 0; i < int(subPackageCount); i++ {
			v, r2 := 0, 0
			v, newcontent, r2 = readPackage(newcontent)
			version += v
			versionPart2 = append(versionPart2, r2)
		}
		content = newcontent
	}
	//part2
	resultPart2 := 0
	switch id {
	case 0:
		resultPart2 = utils.SumInts(versionPart2)
	case 1:
		resultPart2 = utils.ProdInts(versionPart2)
	case 2:
		resultPart2 = utils.MinInts(versionPart2)
	case 3:
		resultPart2 = utils.MaxInts(versionPart2)
	case 5:
		resultPart2 = utils.Greater(versionPart2[0], versionPart2[1])
	case 6:
		resultPart2 = utils.Less(versionPart2[0], versionPart2[1])
	case 7:
		resultPart2 = utils.Equal(versionPart2[0], versionPart2[1])
	}

	return version, content, resultPart2
}
