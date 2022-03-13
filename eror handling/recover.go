package main

import "fmt"

func endApps() {
	message := recover()
	fmt.Println("telah terjadi eror seperti : ", message)
	fmt.Println("aplikasi telah berhenti")

}
func runApps(error bool) {
	defer endApps()
	if error {
		panic("ERROR EROR")
	}
	fmt.Println("aplikasi tetap berjalan")
}
func main() {
	runApps(true)
}
