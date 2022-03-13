package main

import "fmt"

type Pribadi struct {
	ID       int
	Fname    string
	Lname    string
	Alamat   string
	Email    string
	Isactive bool
}

func main() {

	person := Pribadi{1, "joko", "ardi", "sobayan", "joko12@yahoo.com", true}

	//cara 1 menampilkan
	//fmt.Println(display(person))

	//cara kedua
	displayPerson := display(person)

	fmt.Println(displayPerson)

}
func display(person Pribadi) string {
	result := fmt.Sprintf("Name : %s %s, Email: %s", person.Fname, person.Lname, person.Email)
	return result
}
