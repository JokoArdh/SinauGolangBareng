package main

import "fmt"

func main() {

	var statusToko = "open"
	var telur = "soldout"
	var buah = "soldout"

	if statusToko == "open" {
		fmt.Println("saya akan pergi belanja telur dan buah")

		if telur == "soldout" || buah == "soldout" {
			fmt.Println("saya tidak jadi beli apa apa di toko")
		} else if telur == "soldout" {
			fmt.Println("telur habis")
		} else if buah == "soldout" {
			fmt.Println("buah habis")
		}
	} else {
		fmt.Println("toko tutup belanja tidak jadi")
	}
}
