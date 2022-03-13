package main

import "fmt"

func main() {
	defer satu()
	dua()

}
func satu() {
	fmt.Println("satu")
}
func dua() {
	fmt.Println("dua")
}
