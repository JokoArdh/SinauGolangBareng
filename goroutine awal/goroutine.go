package main

import (
	"fmt"
	"time"
)

func main() {

	go say() // inisialisasi goroutine

	time.Sleep(5 * time.Second)
	fmt.Println("ini fungsi utama")
}
func say() {
	fmt.Println("halo joko ardiyanto")
}
