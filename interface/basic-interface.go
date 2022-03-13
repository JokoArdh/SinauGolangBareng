package main

import "fmt"

type sapaan interface {
	sayHello() string
}

type person struct {
	name string
	age  int
}

func (p person) sayHello() string {
	return "Halo my name is " + p.name
}

func main() {

	var john sapaan = person{name: "joko", age: 22}

	fmt.Println(john)
}
