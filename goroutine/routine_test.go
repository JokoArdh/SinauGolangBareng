package gorutin

import (
	"fmt"
	"testing"
	"time"
)

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}
func TestNumber(t *testing.T) {

	for i := 0; i <= 100; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
