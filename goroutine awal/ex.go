package main

import (
	"fmt"
	"time"
)

func main() {
	go CetakNama("pagi")
	go CetakNama("sore")

	fmt.Scanln() //fungsi untuk memberhentikan asynchronous
}
func CetakNama(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(2 * time.Second)
	}
}
