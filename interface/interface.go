package main

import "fmt"

type Luas interface {
	Hitungluas() int
}
type Persegi struct {
	sisi int
}

func (persegi Persegi) Hitungluas() int {
	return persegi.sisi * persegi.sisi
}

func main() {
	bangundatar := Persegi{sisi: 5}
	luas := bangundatar.Hitungluas()
	fmt.Println("luas : ", luas)
}
