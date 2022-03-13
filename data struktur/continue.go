package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue // untuk mengskip jadi semisal bil habis dibagi 2 tidak ditampilkan lagi
		}

		fmt.Println("pelurangan ke -", i)
	}
}
