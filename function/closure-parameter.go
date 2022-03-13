package main

import "fmt"

// function yang menggunakan function sebagai parameternya
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

// function yang digunakan sebagai parameter
func spamFilter(name string) string {
	if name == "Kasar" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("John", spamFilter)
	sayHelloWithFilter("Kasar", spamFilter)
}
