package main

import "fmt"

func main() {

	score := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	var total int

	for _, score := range score { //kenapa memakai for perulangan karena nanti penjumlaham berulang dari scro 100+80 dst
		//rumus total
		total = total + score
	}
	//rata-rata
	length := len(score)                        //mengambil nilai score
	average := float64(total) / float64(length) //rumus dari rata-rata dengan tipe float64 agar nilainya bil bulat

	fmt.Println("terdapat data sbg berikut :", score)
	fmt.Println("=================================================")
	fmt.Println("total dari score :", total)
	fmt.Println("hasil dari rata- rata :", average)

}
