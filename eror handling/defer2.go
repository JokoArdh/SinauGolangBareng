package main

import "fmt"

func main() {

	orderFood("pizza")
	orderFood("burger")

}
func orderFood(menu string) {
	defer fmt.Println("Terimakasih, Please wait !")

	if menu == "pizza" {
		fmt.Print("pilihan tepat !", "")
		fmt.Print("pizza di tempat kami paling enak ", "\n")
		return
	}
	fmt.Println("pesanan anda : ", menu)
}
