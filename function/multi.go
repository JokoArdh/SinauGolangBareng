package main

import "fmt"

func main() {
	//deklarasi
	first, last := intro("joko", "ardiyanto")
	fmt.Println(first, last)

	first1, _ := intro("joko", "ardiyanto")
	fmt.Println(first1)

}

func intro(first string, last string) (string, string) {
	introFirst := "Halo nama depan saya  " + first
	introLast := "Halo nama belakang saya " + last
	return introFirst, introLast
}
