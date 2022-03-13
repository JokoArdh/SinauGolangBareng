package main

import "fmt"

func main() {

	bateraiHp := 65
	statusHp := "hidup"
	wifi := false
	bluetooth := false
	whatsapp := "open"

	if statusHp == "mati" && bateraiHp >= 90 {
		fmt.Println("Hp sedang mati")
	} else if statusHp == "hidup" && whatsapp == "open" {
		if wifi == false && bluetooth == true && whatsapp == "open" {
			fmt.Println("Whatsapp sedang dibuka tapi bluetooth hidup")
		} else {
			fmt.Println("wifi sedang hidup dan sedang membuka whatsapp")
		}
	}
}
