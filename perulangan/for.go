package main

import "fmt"

func main() {

	//for i := 0; i < 5; i++ {
	//	fmt.Println("belajar peulangan ke -", i)
	//}

	//i := 1   (PENGGANGTI WHILE DI GOLANG)
	//for i <= 5 {
	//	fmt.Println("belajar peulangan ke -", i)
	//	i++
	//}

	title := "Golang is the best language" //konsep hampir sama foreach
	for index, letter := range title {
		fmt.Println("index :", index, "letter :", string(letter))
	}
}
