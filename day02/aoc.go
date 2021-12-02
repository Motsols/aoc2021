package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)
type Action struct {
	Direction string
	Distance int
}

func getSolutionPart1(input []Action) int {
	forward, depth := 0, 0

	for _, action := range input {
		if(action.Direction == "forward"){
			forward += action.Distance
		}
		if(action.Direction == "down"){
			depth += action.Distance
		}
		if(action.Direction == "up"){
			depth -= action.Distance
		}
	}

	return forward * depth
}

func getSolutionPart2(input []Action) int {
	forward, depth, aim := 0, 0, 0

	for _, action := range input {
		if(action.Direction == "forward"){
			forward += action.Distance
			depth += aim * action.Distance
		}
		if(action.Direction == "down"){
			aim += action.Distance
		}
		if(action.Direction == "up"){
			aim -= action.Distance
		}
	}

	return forward * depth
}

func parseInput(input string) ([]Action, error) {
	var actions []Action
	lines := strings.Split(input, "\n")
	
	for _, line := range lines {
		var parts = strings.Split(strings.TrimSpace(line), " ")
		distance, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		actions = append(actions, Action{parts[0], distance})
	}

	return actions, nil
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
