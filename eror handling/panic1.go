package main

import "fmt"

func Endapp() {

	message := recover()
	if message != nil {
		fmt.Println("Eror dengan message , ", message)
	}

	fmt.Println("aplikasi selesai")
}

func Runapp(error bool) {
	defer Endapp()

	if error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("aplikasi berjalan")
}

func main() {
	Runapp(true)
	//Runapp(false)
}
