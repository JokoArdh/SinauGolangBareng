package main

import "fmt"

func main() {

	//menggunakan slice karena data bersifat dinamis

	// [	PENULISAN 1]
	//var gamingConsole []string
	//gamingConsole = append(gamingConsole, "mobile legend")
	//gamingConsole = append(gamingConsole, "free fire")
	//gamingConsole = append(gamingConsole, "arena of voler")

	// [PENULISAN II]
	gamingConsole := []string{"mobile legend", "arena of valor", "free fire"}

	fmt.Println(gamingConsole)
}
