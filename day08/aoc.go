package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	numbers, out := parseInput(string(inputBytes))

	part := os.Getenv("part")

	if part == "part2" {
		fmt.Println(getPart2(numbers, out))
	} else {
		fmt.Println(getPart1(out))
	}
}
func parseInput(input string) ([]string, []string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var numbers, out = make([]string, len(lines)), make([]string, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, "|")
		numbers[i] = parts[0]
		out[i] = parts[1]
	}

	return numbers, out
}

func getPart1(outputs []string) int {
	hits := 0
	for _, output := range outputs {
		var stringNumbers = strings.Split(strings.TrimSpace(output), " ")
		for _, num := range stringNumbers {
			switch len(num) {
			case 2, 3, 4, 7:
				hits++
			}
		}
	}
	return hits
}
func getPart2(numbers []string, out []string) int {
	return 1
}
