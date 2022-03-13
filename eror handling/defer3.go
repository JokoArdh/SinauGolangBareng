package main

import "fmt"

func main() {

	number := 3

	if number == 3 {
		fmt.Println("halo 1")
		// defer fmt.Println("halo 3")

		func() {
			defer fmt.Println("halo 3")
		}()
	}

	fmt.Println("halo 2")
}
