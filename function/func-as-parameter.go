package main

import "fmt"

func main() {

	sayHello("joko", spamFilter)
	sayHello("jahat", spamFilter)

}

//funcion yang digunakan sebagai func as parameter
func sayHello(name string, filter func(string) string) {
	filtering := filter(name)
	fmt.Println("haloo syang ku ", filtering)
}

//func yang dijadikan parameter
func spamFilter(name string) string {
	if name == "jahat" {
		return "..."
	} else {
		return name
	}
}
