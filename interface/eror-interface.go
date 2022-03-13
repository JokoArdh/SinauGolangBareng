package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {

	//CONTOH penggunaan erorrs
	//var contohEror error = errors.New("ups.. eror")
	//fmt.Println(contohEror.Error())

	hasil, err := pembagi(100, 10) // jika 10 diganti 0 maka akan muncul pemberitahuan Error
	if err == nil {
		fmt.Println("Hasil = ", hasil)
	} else {
		fmt.Println("Eror =", err.Error())
	}
}
