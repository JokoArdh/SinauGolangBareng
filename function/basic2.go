package main

import "fmt"

func sayHi(fistName, lastName string, age int) {
	fmt.Println("Hello", fistName, lastName, age, "years old")
}

func main() {
	sayHi("joko", "ardi", 27)
}
