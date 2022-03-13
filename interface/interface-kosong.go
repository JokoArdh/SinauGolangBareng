package main

import "fmt"

// Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis	semua tipe data akan menjadi implementasinya
// Interface kosong biasa digunakan untuk menampung tipe data yg dynamic

/*cara membuat interface kosong -> interface{} atau sama dengan
type Apapun interface {
}*/

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func main() {
	//data := Ups(2)
	var data interface{} = Ups(1)
	fmt.Println(data)
}
