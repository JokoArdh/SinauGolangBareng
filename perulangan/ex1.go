package main

import (
	"fmt"
)

func luasSegitiga(alas int, tinggi int) int {
	return alas * tinggi / 26
}

func main() {

	fmt.Println(luasSegitiga(5, 6))

}
