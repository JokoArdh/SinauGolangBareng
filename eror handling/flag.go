package main

import (
	"flag"
	"fmt"
)

func main() {

	var panjang = flag.Int("panjang", 5, "input panjang balok")
	var lebar = flag.Int("lebar", 4, "input lebar balok")
	var tinggi = flag.Int("tinggi", 6, "input tinggi balok")

	volumeBalok := (*panjang) * (*lebar) * (*tinggi)
	fmt.Println(volumeBalok)
}
