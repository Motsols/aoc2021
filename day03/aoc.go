package main; import ("fmt";"io/ioutil";"os";"strconv";"strings");


func getSolutionPart1(input []string) int {
	lineLength := len(input[0])
	var positions = make([][2]int, lineLength)

	for _, line := range input {
		var readings = strings.Split(line, "");
		for x, reading := range readings {
			switch reading {
				case "0": positions[x][0]++
				case "1": positions[x][1]++
			}
		}
	}

	gamma, epsilon := "", ""
	for i := 0; i < lineLength; i++ {
		if positions[i][0] > positions[i][1] {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		} else { 
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		}
	}

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)

	return int(gammaInt) * int(epsilonInt)
}

func getSolutionPart2(input []string) int {
	lineLength := len(input[0])
	oxygenLines := input
	co2Lines := input

	for i := 0; i < lineLength; i++ {
		if(len(oxygenLines) == 1) {
			break
		}

		positionCounts := getPositionCount(oxygenLines, i)

		mostCommon := "1"
		if positionCounts[0] > positionCounts[1] {
			mostCommon = "0"
		}

		oxygenLines = getLines(oxygenLines, mostCommon, i)
	}
	for i := 0; i < lineLength; i++ {
		if(len(co2Lines) == 1) {
			break
		}
		
		positionCounts := getPositionCount(co2Lines, i)

		leastCommon := "0"
		if positionCounts[0] > positionCounts[1] {
			leastCommon = "1"
		}

		co2Lines = getLines(co2Lines, leastCommon, i)
	}

	oxygenInt, _ := strconv.ParseInt(oxygenLines[0], 2, 64)
	co2Int, _ := strconv.ParseInt(co2Lines[0], 2, 64)

	return int(oxygenInt) * int(co2Int)
}

func getPositionCount(input []string, position int) [2]int {
	var positionCounts [2]int

	for _, line := range input {
		var readings = strings.Split(line, "");
		if (readings[position] == "0") {
			positionCounts[0]++;
		} else {
			positionCounts[1]++;
		}
	}

	return positionCounts
}

func getLines(input []string, pickValue string, pickPosition int) []string {
	var remainingLines []string
	
	for _, line := range input {
		var readings = strings.Split(line, "");
		if readings[pickPosition] == pickValue {
			remainingLines = append(remainingLines, line)
		}
	}

	return remainingLines
}

func parseInput(input string) ([]string) {
	lines := strings.Split(input, "\r\n")

	return lines
}

func main() {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("couldn't read input")
	}

	input := parseInput(string(inputBytes))

	fmt.Println("Go")
	part := os.Getenv("part")

	if part == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}