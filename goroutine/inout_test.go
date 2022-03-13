package gorutin

import (
	"fmt"
	"testing"
	"time"
)

//membuat sebuah channel IN & OUT
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)

	channel <- "joko ardiyanto"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestChannelOly(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}
