package main

import "fmt"

func main() {

	//tipe number decimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	//tipe number
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	//tipe string
	var pesan = "halo semua salam kenal"
	fmt.Println(pesan)

	//example
	var name = "John Doe"
	var age = "28"
	var sentence = `halo nama saya "` + name + `" dan berumur "` + age + `"`

	fmt.Println(sentence)
}
