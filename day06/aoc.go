package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func getSolution(input []int, days int) int {
	// dayGroup := 2
	// miniDayz := days/dayGroup // 80/4=20

	for i := 0; i < days; i++ {
		for fish, value := range input {
			if value == 0{
				input[fish] = 6
				input = append(input, 8)
			} else {
				input[fish]--
			}
		}
	}

	// for i := 0; i < miniDayz; i++ {
	// 	for fish, value := range input { // 3, 4,3,1,2
	// 		pregnancy := value-dayGroup
	// 		// fmt.Printf("Pregnancy: %d \r\n", pregnancy)
	// 		if pregnancy <= 0{
	// 			// fmt.Printf("Day %d. Fish %d Pregnancy was %d now %d, will now get 6-pregnancy for a total of %d(6). New fish will have %d(8) \r\n", i*dayGroup, fish, input[fish], pregnancy, 6+pregnancy, 8+pregnancy)
	// 			input[fish] = 6+pregnancy
	// 			input = append(input, 8+pregnancy)
	// 		} else {
	// 			// fmt.Printf("Day %d. Fish %d Not giving birth yet. Pregnancy was %d now %d \r\n", i*dayGroup, fish, input[fish], pregnancy)
	// 			input[fish] = pregnancy
	// 		}
	// 	}
	// }

	return len(input)
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
	inputBytes, err := ioutil.ReadFile("testinput.txt")
	if err != nil {
		panic("couldn't read input")
	}

	input, err := parseInput(string(inputBytes))
	if err != nil {
		panic("couldn't parse input")
	}

	fmt.Println("Go")
	part := os.Getenv("part")

	// fmt.Println(getSolution(input, 256))
	if part == "part2" {
		fmt.Println(getSolution(input, 256))
	} else {
		fmt.Println("Go")
		fmt.Println(getSolution(input, 80))
	}
}
