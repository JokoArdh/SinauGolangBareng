package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	var john = student{"john erik", 23}
	john.sayHello()
}

func (s student) sayHello() {
	fmt.Println("Helo ", s.name)
}
