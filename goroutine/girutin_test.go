package gorutin

import (
	"fmt"
	"testing"
	"time"
)

func word() {
	fmt.Println("halooo semuanya")
}

func TestWord(t *testing.T) {
	go word()
	fmt.Println("OOPS")

	time.Sleep(1 * time.Second)
}
