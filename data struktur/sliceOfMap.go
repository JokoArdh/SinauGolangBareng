package main

import "fmt"

func main() {
	student := []map[string]string{
		{"name": "joko", "score": "B"},
		{"name": "andy", "score": "C"},
		{"name": "putri", "score": "B"},
		{"name": "enlo", "score": "E"},
	}

	//fmt.Println(student)

	//perulangan untuk cetak satu satu
	for _, student := range student {
		fmt.Println(student["name"], "_", student["score"])
	}
}
