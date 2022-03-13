package main

import "fmt"

func main() {
	//age := 9

	//if age > 10 {
	//	fmt.Println("udah boleh pilih pasangan pendamping hidup kan udah dewasa")
	//} else {
	//	fmt.Println("tidak boleh")
	//}

	score := 65
	var grade string

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score <= 70 {
		grade = "C"
	} else {
		grade = "NULL"
	}
	fmt.Println(grade)
}
