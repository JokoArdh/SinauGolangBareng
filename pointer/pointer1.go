package main

import "fmt"

func main() {

	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   : ", numberA)
	fmt.Println("numberA (address) : ", &numberA)
	fmt.Println("numberB (value)   : ", *numberB)
	fmt.Println("numberB (address) : ", numberB)

	fmt.Println("==========PERUBAHAN==========")

	numberA = 5 //perubahan nilai yang mempengaruhi keseluruhan bvariabel pojniter

	fmt.Println("numberA (value)   : ", numberA)
	fmt.Println("numberA (address) : ", &numberA)
	fmt.Println("numberB (value)   : ", *numberB)
	fmt.Println("numberB (address) : ", numberB)

}
