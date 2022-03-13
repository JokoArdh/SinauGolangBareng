package main

import (
	"fmt"
	"sync"
)

func printText(text string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go printText("Halo", &wg)

	wg.Add(1)
	go printText("Dunia", &wg)

	wg.Wait()
}
