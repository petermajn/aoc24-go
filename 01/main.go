package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//go:embed example.txt
var example string
var input string

func main() {
	part1(example)
}

func part1(input string) int {
	fmt.Println(example)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	fmt.Println(lines)

	var leftNumbers []int
	var rightNumbers []int

	for _, line := range lines {
		fmt.Println(line)
		lineSplit := strings.Split(line, "   ")
		leftNum, err1 := strconv.Atoi(lineSplit[0])
		rightNum, err2 := strconv.Atoi(lineSplit[1])

		if err1 != nil || err2 != nil {
			fmt.Printf("Error parsing numbers on line: \"%s\"\n", line)
			continue
		}
		leftNumbers = append(leftNumbers, leftNum)
		rightNumbers = append(rightNumbers, rightNum)

	}
	fmt.Println("Left numbers:", leftNumbers)
	fmt.Println("Right numbers:", rightNumbers)
	return 1

}
