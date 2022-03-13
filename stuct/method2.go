package main

import "fmt"

type User struct {
	ID       int
	Fname    string
	Lname    string
	Email    string
	Isactive bool
}

type Grup struct {
	Name        string
	Admin       User
	Users       []User
	Isavailable bool
}

//method
func (user User) display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.Fname, user.Lname, user.Fname)
}

func main() {

	user1 := User{1, "joko", "ardiyanto", "joko@yahoo.com", true}
	result := user1.display()
	fmt.Println(result)

	user2 := User{2, "erna", "wulandari", "ernacantik@yahoo.com", true}
	fmt.Println(user2.display()) //memanggil sebuah user 2

}
