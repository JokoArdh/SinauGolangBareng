package main

import "fmt"

func main() {

	//menentukan modulus atau mencetak angka yang indexnya genap solusi 1
	title := "Golang is the best language"
	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("index :", index, "letter :", string(letter))
		}

	}

}
