package main

import "fmt"

type Blachat func(string) bool

//kode anonymous function
func loginUser(name string, blackhat Blachat) {
	if blackhat(name) {
		fmt.Println("access  is blocked", name)
	} else {
		fmt.Println("access succes", name)
	}
}

func main() {

	blackhat := func(name string) bool {
		return name == "root"
	}

	loginUser("root", blackhat)
	loginUser("joko", blackhat)

	loginUser("root", func(name string) bool {
		return name == "root"
	})
	loginUser("joko", func(name string) bool {
		return name == "root"
	})

}
