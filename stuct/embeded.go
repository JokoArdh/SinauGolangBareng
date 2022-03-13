package main

import "fmt"

type person struct {
	name string
	age  int
}
type student struct {
	grade  int
	person // menampung struct pertama yg bernama person
}

func main() {

	//example 1
	var john = student{}
	john.name = "john"
	john.age = 21
	john.grade = 4

	fmt.Println("===================PERTAMA===============")
	fmt.Println("Name  : ", john.name)
	fmt.Println("Age   : ", john.age)
	fmt.Println("Umur  : ", john.person.age)
	fmt.Println("Grade : ", john.grade)

	//example 2

	var doeData = person{name: "Doel", age: 23}
	var doe = student{person: doeData, grade: 3}

	fmt.Println("===================KEDUA===============")
	fmt.Println("Name  : ", doe.name)
	fmt.Println("Age   : ", doe.age)
	fmt.Println("Grade : ", doe.grade)

}
