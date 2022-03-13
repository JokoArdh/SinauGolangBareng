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

func main() {

	user1 := User{1, "joko", "ardiyanto", "joko@yahoo.com", true}
	user2 := User{2, "erna", "wulandari", "ernacantik@yahoo.com", true}

	users := []User{user1, user2} // katrena di struct grub users itu slice  []

	grup := Grup{"gamers", user1, users, true}

	dispalyGrup(grup)
}

//membuat function display grub
func dispalyGrup(grup Grup) {
	// ketika mau memangil isi dari sebuah variabel input maka harus printf karena memakai karakter %s dan %d
	fmt.Printf("Name : %s ", grup.Name)
	fmt.Println("")
	fmt.Printf("Member Count : %d", len(grup.Users)) //%d untuk nilai integer
	fmt.Println("")

	//ketika ingin menampilkan sebuah looping untuk nama yg ada di grub
	fmt.Println("==============User List=================")
	for _, user := range grup.Users {
		fmt.Println(user.Fname)
	}
}
