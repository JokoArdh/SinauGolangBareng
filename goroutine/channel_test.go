package gorutin

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {

	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "joko ardiyanto"
		fmt.Println("selesai mengirim data ke channel")
	}()

	//mengambil channel
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}

func GiveMeAnswer(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "programmer universitas boyolali"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeAnswer(channel)

	//mengambil channel
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}
