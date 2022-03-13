package main

import "fmt"

func main() {

	result := addNumber(10, 7)
	fmt.Println(result)

}
func addNumber(number int, numberTwo int) int {
	result := number + numberTwo
	return result

	// atau
	// return number + numberTwo
}
