package gorutin

import (
	"fmt"
	"testing"
)

func TestBuffer(t *testing.T) {
	channel := make(chan string, 3) //menggunakan channel dengan buffer maximal 3 data
	defer close(channel)

	channel <- "joko"
	channel <- "universitas boyolali"
	channel <- "fakultas ilmu kommputer"
	//dengan kualioifikasi data channel sesuai yaitu 3

	//mengambil data channel
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("selesai")
}
