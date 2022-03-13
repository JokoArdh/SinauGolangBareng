package main

import "fmt"

func main() {
	var total = sum(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)

	fmt.Println(total)

}
func sum(number ...int) int {
	var total int = 0
	for _, numbers := range number {
		total += numbers
	}
	return total

}
