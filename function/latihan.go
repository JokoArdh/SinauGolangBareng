package main

import "fmt"

func main() {

	score := []int{10, 5, 8, 9, 7}
	var total = Mean(score...)

	fmt.Println(total)

}

func Mean(number ...int) int {
	var total int = 0

	for _, score := range number {
		total += score
	}

	return total
}

//varioadic function
