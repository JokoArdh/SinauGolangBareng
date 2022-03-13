package main

import (
	"errors"
	"fmt"
	"strings"
)

//membuat custom eror sendiri

func validator(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("tidak boleh kosong")
	}
	return true, nil
}

func main() {

	var name string
	fmt.Print("Masukan angka : ")
	fmt.Scanln(&name)

	if valid, err := validator(name); valid {
		fmt.Println("halo ", name)
	} else {
		fmt.Println(err.Error())
	}
}
