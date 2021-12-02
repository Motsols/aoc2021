package main; import ("fmt";"io/ioutil";"os";"strconv";"strings"); 

func main() { inputBytes, _ := ioutil.ReadFile("input.txt"); input := parseInput(string(inputBytes)); fmt.Println("Go"); part := os.Getenv("part"); if part == "part2" { fmt.Println(getPart2(input)) } else { fmt.Println(getPart1(input)); } };
func parseInput(input string) ([]int) { var ints []int; lines := strings.Split(strings.TrimSpace(input), "\r\n"); for _, line := range lines { i, _ := strconv.Atoi(line); ints = append(ints, i); }; return ints }; 

func getPart1(input []int) int { increases, lastInput := 0, 99999; for _, value := range input { if value > lastInput{ increases++ }; lastInput = value }; return increases }; 
func getPart2(input []int) int { increases := 0; lastSum := 0; for i, _ := range input { newSum := 0; if i > 1 { newSum = input[i] + input[i-1] + input[i-2]; if(newSum > lastSum){ increases++; }; }; lastSum = newSum; }; return increases };