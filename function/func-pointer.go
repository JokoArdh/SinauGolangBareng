package main

import "fmt"

func tambahKeranjangBelanja(belanjaan string, keranjang *[]string) {
	*keranjang = append(*keranjang, belanjaan)
}

func main() {

	keranjangBelanjaan := []string{}

	tambahKeranjangBelanja("sayur", &keranjangBelanjaan)
	tambahKeranjangBelanja("buah", &keranjangBelanjaan)
	tambahKeranjangBelanja("telur", &keranjangBelanjaan)
	tambahKeranjangBelanja("beras", &keranjangBelanjaan)
	tambahKeranjangBelanja("minyak", &keranjangBelanjaan)

	fmt.Println(keranjangBelanjaan)

}
