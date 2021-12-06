package main
import ( "fmt"; "os" )

func getSolution(startingFish []int, days int) int {
	fishPregnancies := map[int]int{}
	longDays := days/4
	zero,one,two,three,four,five,six,seven,eigth := 0,0,0,0,0,0,0,0,0

	for _ , daysLeft := range startingFish {
		fishPregnancies[daysLeft]++
	}

	for i := 0; i < longDays; i++ {
		zero,one,two,three,four,five,six,seven,eigth = fishPregnancies[0], fishPregnancies[1], fishPregnancies[2], fishPregnancies[3], fishPregnancies[4], fishPregnancies[5], fishPregnancies[6], fishPregnancies[7], fishPregnancies[8]
		fishPregnancies[0] = four
		fishPregnancies[1] = five
		fishPregnancies[2] = six
		fishPregnancies[3] = zero + seven // parents + newborn
		fishPregnancies[4] = one + eigth // parents + newborn
		fishPregnancies[5] = zero + two // parents + newborn
		fishPregnancies[6] = one + three // parents + newborn
		fishPregnancies[7] = two
		fishPregnancies[8] = three
	}

	totalFishes := 0
	for _, numberOfFishWithXDaysLeft := range fishPregnancies {
		totalFishes += numberOfFishWithXDaysLeft
	}

	return totalFishes
}

func main() {
	// testInput := []int{3,4,3,1,2}
	input := []int{3,4,3,1,2,1,5,1,1,1,1,4,1,2,1,1,2,1,1,1,3,4,4,4,1,3,2,1,3,4,1,1,3,4,2,5,5,3,3,3,5,1,4,1,2,3,1,1,1,4,1,4,1,5,3,3,1,4,1,5,1,2,2,1,1,5,5,2,5,1,1,1,1,3,1,4,1,1,1,4,1,1,1,5,2,3,5,3,4,1,1,1,1,1,2,2,1,1,1,1,1,1,5,5,1,3,3,1,2,1,3,1,5,1,1,4,1,1,2,4,1,5,1,1,3,3,3,4,2,4,1,1,5,1,1,1,1,4,4,1,1,1,3,1,1,2,1,3,1,1,1,1,5,3,3,2,2,1,4,3,3,2,1,3,3,1,2,5,1,3,5,2,2,1,1,1,1,5,1,2,1,1,3,5,4,2,3,1,1,1,4,1,3,2,1,5,4,5,1,4,5,1,3,3,5,1,2,1,1,3,3,1,5,3,1,1,1,3,2,5,5,1,1,4,2,1,2,1,1,5,5,1,4,1,1,3,1,5,2,5,3,1,5,2,2,1,1,5,1,5,1,2,1,3,1,1,1,2,3,2,1,4,1,1,1,1,5,4,1,4,5,1,4,3,4,1,1,1,1,2,5,4,1,1,3,1,2,1,1,2,1,1,1,2,1,1,1,1,1,4}
	
	part := os.Getenv("part")

	if part == "part2" { fmt.Println(getSolution(input, 256))
	} else { fmt.Println(getSolution(input, 80)) }
}