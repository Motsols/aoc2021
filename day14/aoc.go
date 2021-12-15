package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	template, insertions := parseInput(string(inputBytes))

	part := os.Getenv("part")

	if part == "part2" {
		fmt.Println(getAnswer(40, template, insertions))
	} else {
		fmt.Println(getAnswer(40, template, insertions))
	}
}
func parseInput(input string) (map[string]uint64, map[string]string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var templatePairs, insertions = make(map[string]uint64), make(map[string]string)

	for i, line := range lines {
		if i == 0 {
			chars := strings.Split(line, "")
			for j, char := range chars {
				if j != len(chars)-1 {
					templatePairs[char+chars[j+1]] += 1
				}
			}
		}
		if i >= 2 {
			parts := strings.Split(line, " -> ")
			insertions[parts[0]] = parts[1]
		}
	}

	return templatePairs, insertions
}

func getAnswer(steps int, templatePairs map[string]uint64, insertions map[string]string) uint64 {
	letters := make(map[string]uint64)

	for pair, _ := range templatePairs {
		parts := strings.Split(pair, "")
		if len(letters) == 0 {
			letters[parts[0]]++
		}
		letters[parts[1]]++
	}

	for i := 0; i < steps; i++ {
		newPairs := make(map[string]uint64)

		for key, count := range templatePairs {
			value, _ := insertions[key]

			parts := strings.Split(key, "")
			newPairs[parts[0]+value] += uint64(count)
			newPairs[value+parts[1]] += uint64(count)
			letters[value] += count
		}

		templatePairs = newPairs
	}

	leastOccurrences, mostOccurrences := ^uint64(0), uint64(0)
	for _, uses := range letters {
		if uses < leastOccurrences {
			leastOccurrences = uses
		} else if uses > mostOccurrences {
			mostOccurrences = uses
		}
	}

	return mostOccurrences - leastOccurrences
}
