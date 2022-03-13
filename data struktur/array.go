package main

import "fmt"

func main() {

	//var languages [5]string (pengisian tdk langsung)

	//languages[0] = "Go"
	//languages[1] = "Js"
	//languages[2] = "Py"
	//languages[3] = "PHP"
	//languages[4] = "C++"

	//languages := [5]string{"python,javascript", "Golang", "php", "c++"} //pngisian langsung

	//cara 3
	languages := [...]string{
		"python",
		"kotlin",
		"golang",
		"javascript",
		"C#",
	}

	//fmt.Println(languages)
	//length := len(languages)
	//fmt.Println(length)

	//menggunsksn perulangan array
	for index, lang := range languages {
		fmt.Println("index :", index, "languages :", lang)
	}
}
