package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func getSolutionPart1(input []int) int {
	increases := 0
	lastInput := 99999

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

func parseInput(input string) ([]int, error) {
	var ints []int

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		ints = append(ints, i)
	}

	return ints, nil
}

func main() {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("couldn't read input")
	}

	input, err := parseInput(string(inputBytes))
	if err != nil {
		panic("couldn't parse input")
	}

	fmt.Println("Go")
	part := os.Getenv("part")

	if part == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}
