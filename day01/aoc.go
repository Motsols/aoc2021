package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func getSolutionPart1(input []int) int {
	increases, lastInput := 0, 99999

	for _, value := range input {
		if value > lastInput{
			increases++
		}
		lastInput = value
	}

	return increases
}

func getSolutionPart2(input []int) int {
	var windows []int
	sum := 0

	for i, _ := range input {
		sum = 0
		if i > 1{
			sum = input[i] + input[i-1] + input[i-2]
			windows = append(windows, sum)
		}
	}

	return getSolutionPart1(windows)
}

func parseInput(input string) ([]int) {
	var ints []int
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		i, _ := strconv.Atoi(line)
		ints = append(ints, i)
	}

	return ints
}

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	input := parseInput(string(inputBytes))
	fmt.Println("Go")
	part := os.Getenv("part")

	if part == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}