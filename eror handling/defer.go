package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}
func runApps() {
	defer logging()
	fmt.Println("Eun applikasi")
}
func main() {
	runApps()
}
