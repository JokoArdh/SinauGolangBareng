package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a = 3 // ubah lebih dari 1

	switch a {
	case 1:
		fmt.Println(strconv.Itoa(a) + "st")
	case 2:
		fmt.Println(strconv.Itoa(a) + "sd")
	case 3:
		fmt.Println(strconv.Itoa(a) + "rd")
		fallthrough
	case 4, 5, 6, 7:
		fmt.Println(strconv.Itoa(a) + "th")
	default:
		fmt.Println("not winner")
	}
}
