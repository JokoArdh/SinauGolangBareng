package main

import "fmt"

func main() {

	luas, keliling := penjumlahan(10, 6)
	fmt.Println(luas, keliling)
}

func penjumlahan(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}
