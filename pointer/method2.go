package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

//membuat addgame (wajib diisi dengan *) karena mengubah internal propresties
func (gamer *Gamer) addGame(game string) {
	gamer.Games = append(gamer.Games, game)
}

func main() {
	gamer := Gamer{Name: "joko ardi"}

	gamer.addGame("mobile legend")
	gamer.addGame("fire fire")
	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
