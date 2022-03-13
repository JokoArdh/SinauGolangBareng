package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "1990-03-20")
	fmt.Println(parse)
	fmt.Println()

	fmt.Println("start")
	time.Sleep(time.Second * 4)
	fmt.Println("after 4 seconds")
}
