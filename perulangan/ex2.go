package main

import (
	"fmt"
)

func main() {
	var str = "Golang"

	jumlah := 0

	for _, char := range str {
		fmt.Println(string(char))
	}

	for i := 0; i <= 5; i++ {
		jumlah += 1
	}
	fmt.Println(jumlah)
}
