package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func getSolution(startingFish []int, days int) int {
	fishPregnancies := map[int]int{}
	births := 0

	for _ , daysLeft := range startingFish {
		fishPregnancies[daysLeft]++
	}

	for i := 0; i < days; i++ {
		births = fishPregnancies[0]
		fishPregnancies[0] = fishPregnancies[1]
		fishPregnancies[1] = fishPregnancies[2]
		fishPregnancies[2] = fishPregnancies[3]
		fishPregnancies[3] = fishPregnancies[4]
		fishPregnancies[4] = fishPregnancies[5]
		fishPregnancies[5] = fishPregnancies[6]
		fishPregnancies[6] = fishPregnancies[7] + births // parent immediately gets pregnant again. All teenagers become pregnant on the same day. Maybe with the same father (drama)
		fishPregnancies[7] = fishPregnancies[8] // Teenager day 2
		fishPregnancies[8] = births // Childhood day 1
	}

	totalFishes := 0
	for _, numberOfFishWithXDaysLeft := range fishPregnancies {
		totalFishes += numberOfFishWithXDaysLeft
	}

	return totalFishes
}

func parseInput(input string) ([]int, error) {
	var ints []int

	lines := strings.Split(strings.TrimSpace(input), ",")

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
		fmt.Println(getSolution(input, 256))
	} else {
		fmt.Println(getSolution(input, 80))
	}
}
