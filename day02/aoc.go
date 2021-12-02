package main; import ("fmt"; "io/ioutil";"os";"strconv";"strings")

func getSolutionPart1(input []string) int {
	forward, depth := 0, 0

	for _, action := range input {
		var command = strings.Split(strings.TrimSpace(action), " ")
		distance, _ := strconv.Atoi(command[1])
		switch command[0] {
			case "forward": forward += distance
			case "down": depth += distance
			case "up": depth -= distance
		}
	}

	return forward * depth
}

func getSolutionPart2(input []string) int { 
	forward, depth, aim := 0, 0, 0; 

	for _, action := range input { 
		var command = strings.Split(strings.TrimSpace(action), " ")
		distance, _ := strconv.Atoi(command[1])
		switch command[0] {
			case "forward": forward += distance; depth += aim * distance;
			case "down": aim += distance
			case "up": aim -= distance
		}
	} 
		
	return forward * depth
}

func parseInput(input string) ([]string) {
	var actions []string; lines := strings.Split(input, "\n")
	for _, line := range lines { actions = append(actions, line) }

	return actions
}

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt"); input := parseInput(string(inputBytes)); fmt.Println("Go"); part := os.Getenv("part")

	if part == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}