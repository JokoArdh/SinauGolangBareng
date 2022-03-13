package main

import (
	"fmt"
)

func random() interface{} {
	return "Ups!"
}

func main() {
	var result interface{} = random()
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	//SWITCH ke type assertions
	switch value := result.(type) {
	case int:
		fmt.Println("value ", value)
	case string:
		fmt.Println("value", value)
	default:
		fmt.Println("Unknow")
	}
}
