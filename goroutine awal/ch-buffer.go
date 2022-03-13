package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 6
	ch <- 7
	ch <- 5

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}




















