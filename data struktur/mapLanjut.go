package main

import "fmt"

func main() {
	myMap := map[string]string{
		"ruby": "is abeatuful",
		"go":   "is super fast",
		"java": "is more difficult",
	}

	for key, value := range myMap {
		fmt.Println("key :", key, "value :", value)
	}

	//menghapus map
	fmt.Println("=====================")

	delete(myMap, "ruby") // fungsi delete untuk menghapus
	//lagi kita lakukan perulangan
	for key, value := range myMap {
		fmt.Println("key :", key, "isi :", value)
	}
}
