package main

import "fmt"

func main() {
	number := 2 //ketika nilai 3 tadi giganti 5 maka akan cetak default

	switch number {
	case 1:
		fmt.Println("satu")
	case 2:
		fmt.Println("dua")
	case 3:
		fmt.Println("toga")
	default:
		fmt.Println("DEPAULT")
	}
}
