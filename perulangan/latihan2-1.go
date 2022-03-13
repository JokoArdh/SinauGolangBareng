package main

import "fmt"

func main() {
	//program nilai lebih baik atau lebih besar

	score := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	var goodScore []int

	for _, score := range score {
		if score >= 90 {
			goodScore = append(goodScore, score)
		}
	}
	fmt.Println("Nilai terbaik :", goodScore)

}
