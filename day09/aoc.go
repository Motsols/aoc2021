package main;
import ("fmt"; "io/ioutil";"strconv";"os";"strings";"sort")

const (
	rowCount = 100
	columnCount = 100
)

var (
	basinSizes []int
	positionsInBasins = make(map[string]bool)
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
		for j, value := range row {
			if isBasin(i, j, value, heightLines) {
				riskLevel += 1 + value
			}
		}
	}
	
	return riskLevel
}

func getSolutionPart2(heightLines [][columnCount]int) int { 
	for i, row := range heightLines {
		for j, value := range row {
			if isBasin(i, j, value, heightLines) {
				clearBasinsMap()
				getBasinSize(i, j, value, heightLines)
				
				if len(positionsInBasins) > 0 {
					basinSizes = append(basinSizes, len(positionsInBasins) +1)
				}
			}
		}
	}

	sort.Ints(basinSizes)
	numOfBasins := len(basinSizes)
	return basinSizes[numOfBasins-1] * basinSizes[numOfBasins-2] * basinSizes[numOfBasins-3]
}

func isBasin(i int, j int, value int, heightLines [][columnCount]int) bool {
	return left(i, j, value, heightLines, false) && right(i, j, value, heightLines, false) && above(i, j, value, heightLines, false) && below(i, j, value, heightLines, false)
}

func left(i int, j int, height int, heightLines [][columnCount]int, withBasins bool) bool {
	return compare(i, j-1, height, heightLines, withBasins)
}
func right(i int, j int, height int, heightLines [][columnCount]int, withBasins bool) bool {
	return compare(i, j+1, height, heightLines, withBasins)
}
func above(i int, j int, height int, heightLines [][columnCount]int, withBasins bool) bool {
	return compare(i-1, j, height, heightLines, withBasins)
}
func below(i int, j int, height int, heightLines [][columnCount]int, withBasins bool) bool {
	return compare(i+1, j, height, heightLines, withBasins)
}

func compare(i int, j int, height int, heightLines [][columnCount]int, withBasins bool) bool {
	if i < 0 || j < 0 || i > rowCount -1 || j > columnCount -1 {
		if withBasins {
			return false
		} else {
			return true
		}
	}

	if withBasins && heightLines[i][j] > height && heightLines[i][j] != 9 {
		positionString := strconv.Itoa(i) + strconv.Itoa(j)
		_, found := positionsInBasins[positionString]
		
		if found == false {
			positionsInBasins[positionString] = true
		}
		
		return true
	} else if !withBasins && height < heightLines[i][j] {
		return true
	}

	return false
}

func getBasinSize(i int, j int, height int, heightLines [][columnCount]int) {
		if !positionsInBasins[strconv.Itoa(i) + strconv.Itoa(j-1)] && left(i, j, height, heightLines, true) {
			getBasinSize(i, j-1, height, heightLines)
		}
		if !positionsInBasins[strconv.Itoa(i) + strconv.Itoa(j+1)] && right(i, j, height, heightLines, true) {
			getBasinSize(i, j+1, height, heightLines)
		}
		if !positionsInBasins[strconv.Itoa(i-1) + strconv.Itoa(j)] && above(i, j, height, heightLines, true) {
			getBasinSize(i-1, j, height, heightLines)
		}
		if !positionsInBasins[strconv.Itoa(i+1) + strconv.Itoa(j)] && below(i, j, height, heightLines, true) {
			getBasinSize(i+1, j, height, heightLines)
		}
}

func clearBasinsMap() {
	for key, _ := range positionsInBasins {
		delete(positionsInBasins, key)
	}
}