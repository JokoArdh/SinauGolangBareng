package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {

	var joko student
	joko.name = "joko ardiyanto"
	joko.grade = 3

	fmt.Println("Nama : ", joko.name)
	fmt.Println("Grade : ", joko.grade)
}
