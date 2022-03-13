package main

import "fmt"

func main() {

	angka := 20

	var angkaBaru *int = &angka

	*angkaBaru += 5
	angka = 20 - *angkaBaru
	angka += 3
	*angkaBaru += 20

	fmt.Println(angka)
}
