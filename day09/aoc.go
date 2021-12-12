package main;
import ("fmt"; "io/ioutil";"strconv";"os";"strings")

const (
	rowCount = 100
	columnCount = 100
)

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt");
	input := parseInput(string(inputBytes)); 
	part := os.Getenv("part")

	if part == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}

func parseInput(input string) [][columnCount]int {
	lines := strings.Split(input, "\r\n")
	// rowCount = len(lines) - 1
	// columnCount = len(lines[0]) - 1
	// fmt.Printf("rows %d and %d, cols %d and %d \r\n", len(lines), rowCount, len(lines[0]), columnCount)
	heightLines := make([][columnCount]int, rowCount)
	
	for i, line := range lines { 
		var cols = strings.Split(line, "")
		for j, col := range cols {
			height, _ := strconv.Atoi(col)
			heightLines[i][j] = height
		}
	}

	return heightLines
}

func getSolutionPart1(heightLines [][columnCount]int) int {
	riskLevel := 0
	for i, row := range heightLines {
		// fmt.Printf("iteration %d, line %d \r\n", i, row)
		// var heights = strings.Split(row, "")
		for j, value := range row {
			// fmt.Printf("row %d id '%s', col %d, val %s \r\n", i, row, j, value)
			
			if left(i, j, value, heightLines) && right(i, j, value, heightLines) && above(i, j, value, heightLines) && below(i, j, value, heightLines) {
				riskLevel += 1 + value
			}
		}
	}
	
	return riskLevel
}

func left(i int, j int, height int, heightLines [][columnCount]int) bool {
	return compare(i, j-1, height, heightLines)
}
func right(i int, j int, height int, heightLines [][columnCount]int) bool {
	return compare(i, j+1, height, heightLines)
}
func above(i int, j int, height int, heightLines [][columnCount]int) bool {
	return compare(i-1, j, height, heightLines)
}
func below(i int, j int, height int, heightLines [][columnCount]int) bool {
	return compare(i+1, j, height, heightLines)
}

func compare(i int, j int, height int, heightLines [][columnCount]int) bool {
	if i < 0 || j < 0 || i > rowCount -1 || j > columnCount -1 {
		return true
	}

	if height < heightLines[i][j] {
		return true
	}

	return false
}

// func checkLeft(i int, j int, heights int) bool {
// 	// func checkLeft(height int, i int, j int, rows int, lineLength int) bool {
// 	// if i < 0 || j < 0 || i == rows || j == lineLength {
// 	// 	return false
// 	// }
// 	if j == 0 {
// 		return false
// 	}
	

	
// 	return false
// }

func getSolutionPart2(heightLines [][columnCount]int) int { 
	// var basins []int
	// for i, row := range heightLines {
	// 	// fmt.Printf("iteration %d, line %d \r\n", i, row)
	// 	// var heights = strings.Split(row, "")
	// 	for j, value := range row {
	// 		// fmt.Printf("row %d id '%s', col %d, val %s \r\n", i, row, j, value)
			
	// 		if left(i, j, value, heightLines) && right(i, j, value, heightLines) && above(i, j, value, heightLines) && below(i, j, value, heightLines) {
	// 			getBasinSize(i, j, value, heightLines)
	// 			size = 1
	// 			basins = append(basins, 1)
	// 		}
	// 	}
	// }
	
	return 1
}

// func getBasinSize(i int, j int, height int, heightLines [][columnCount]int) int {

// }