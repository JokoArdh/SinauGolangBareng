package main

import "fmt"

func main() {
	//memanggil function di function utama
	printMyResult()
	//memanggil function berparameter
	myWord("halo selamat datang")
}

//membuat function
func printMyResult() {
	fmt.Println("belajar Go")
}

//function berparameter
func myWord(kata string) {
	fmt.Println(kata)
}
