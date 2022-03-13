package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input string
	fmt.Print("Masukan angka : ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is benar")
	} else {
		fmt.Println(input, "is bukan angka")
		fmt.Println(err.Error())
	}

}
