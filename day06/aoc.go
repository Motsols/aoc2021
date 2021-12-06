package main
import ( "fmt"; "os" )

func getSolution(startingFish []int, days int) int {
	var fishes [9]int
	longDays := days/8
	zero,one,two,three,four,five,six,seven,eigth := 0,0,0,0,0,0,0,0,0

	for _ , daysLeft := range startingFish {
		fishes[daysLeft]++
	}

	for i := 0; i < longDays; i++ {
		zero,one,two,three,four,five,six,seven,eigth = fishes[0], fishes[1], fishes[2], fishes[3], fishes[4], fishes[5], fishes[6], fishes[7], fishes[8]
		fishes[0] = eigth + one // parent + newborn
		fishes[1] = zero + two // parent + newborn
		fishes[2] = one + three // parent + newborn
		fishes[3] = two + four // parent + newborn
		fishes[4] = three + five // parent + newborn
		fishes[5] = four + six // parent + newborn
		fishes[6] = five + seven + zero // parent + parent + newborn
		fishes[7] = six // parent + newborn
		fishes[8] = seven + zero // parent + newborn
	}

	totalFishes := 0
	for _, fish := range fishes {
		totalFishes += fish
	}

	return totalFishes
}

func main() {
	input := []int{3,4,3,1,2,1,5,1,1,1,1,4,1,2,1,1,2,1,1,1,3,4,4,4,1,3,2,1,3,4,1,1,3,4,2,5,5,3,3,3,5,1,4,1,2,3,1,1,1,4,1,4,1,5,3,3,1,4,1,5,1,2,2,1,1,5,5,2,5,1,1,1,1,3,1,4,1,1,1,4,1,1,1,5,2,3,5,3,4,1,1,1,1,1,2,2,1,1,1,1,1,1,5,5,1,3,3,1,2,1,3,1,5,1,1,4,1,1,2,4,1,5,1,1,3,3,3,4,2,4,1,1,5,1,1,1,1,4,4,1,1,1,3,1,1,2,1,3,1,1,1,1,5,3,3,2,2,1,4,3,3,2,1,3,3,1,2,5,1,3,5,2,2,1,1,1,1,5,1,2,1,1,3,5,4,2,3,1,1,1,4,1,3,2,1,5,4,5,1,4,5,1,3,3,5,1,2,1,1,3,3,1,5,3,1,1,1,3,2,5,5,1,1,4,2,1,2,1,1,5,5,1,4,1,1,3,1,5,2,5,3,1,5,2,2,1,1,5,1,5,1,2,1,3,1,1,1,2,3,2,1,4,1,1,1,1,5,4,1,4,5,1,4,3,4,1,1,1,1,2,5,4,1,1,3,1,2,1,1,2,1,1,1,2,1,1,1,1,1,4}
	
	fmt.Println("Go")
	part := os.Getenv("part")

	if part == "part2" { fmt.Println(getSolution(input, 256))
	} else { fmt.Println(getSolution(input, 80)) }
}