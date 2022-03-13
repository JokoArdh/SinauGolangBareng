package gorutin

import (
	"fmt"
	"testing"
	"time"
)

func GiveAnswer(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "programmer universitas boyolali"
}
func TestSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveAnswer(channel1)
	go GiveAnswer(channel2)

	//mengambil data dari select channel secara manual
	// select {
	// case data := <-channel1:
	// 	fmt.Println("data dari channel 1", data)
	// case data := <-channel2:
	// 	fmt.Println("data dari channel 2", data)
	// }
	// select {
	// case data := <-channel1:
	// 	fmt.Println("data dari channel 1", data)
	// case data := <-channel2:
	// 	fmt.Println("data dari channel 2", data)
	// }

	//dengan perulamgam
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}

}

//MENGGUNAKAN DEFAULT SELECT CHANNEL
func TestDefaultSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveAnswer(channel1)
	go GiveAnswer(channel2)

	//dengan perulamgam
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2", data)
			counter++

		default:
			fmt.Println("Mennunggu data ..")
		}

		if counter == 2 {
			break
		}
	}

}
