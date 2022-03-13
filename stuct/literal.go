// penulisan struct berbagai variasi seperti dibawah ini
package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {

	//pertama
	var john = student{}
	john.name = "mahendra"
	john.grade = 2

	//kedua dengan property berurutan
	var doe = student{"doe arial", 4}

	//ketiga dengan properti tdk berurutan
	var jack = student{name: "jackiyen", grade: 4}

	fmt.Println("Student 1 : ", john.name)
	fmt.Println("Student 2 : ", doe.name)
	fmt.Println("Student 3 : ", jack.name)
}
