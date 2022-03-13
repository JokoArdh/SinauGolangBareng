package main

import "fmt"

func endApps() {
	fmt.Println("End Application")
}
func runApps(error bool) {
	defer endApps()
	if error {
		panic("EROR")
	}
}
func main() {
	runApps(true)
}
