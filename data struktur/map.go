package main

import "fmt"

func main() {

	var myMap map[string]int
	myMap = map[string]int{}

	myMap["javascript"] = 9
	myMap["golang"] = 10
	myMap["typescript"] = 7

	fmt.Println(myMap["golang"]) //jika memanggil index
	//fmt.Println(myMap)
}
