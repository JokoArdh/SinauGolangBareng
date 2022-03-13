//pointer dengan parameter
package main

import "fmt"

func main() {

	var number = 4
	fmt.Println("nilai awal : ", number) //ssebelum diubah

	change(&number, 10)
	fmt.Println("nilai berubah : ", number)
}
func change(original *int, value int) {
	*original = value
}
