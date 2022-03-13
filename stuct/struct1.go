package main

import "fmt"

type User struct {
	ID        int
	Fisrtname string
	Lastname  string
	Email     string
	Isactive  bool
}

func main() {

	//penulisan cara 1
	//	user := User{}
	//	user.ID = 1
	//	user.Fisrtname = "joko"
	//	user.Lastname = "ardiyanto"
	//	user.Email = "jokoganteng@yahoo.co.id"
	//	user.Isactive = true

	//	user2 := User{}
	//	user2.ID = 2
	//	user2.Fisrtname = "fly"
	//	user2.Lastname = "ngeflay"
	//	user2.Email = "ngeflay@yahoo.co.id"
	//	user2.Isactive = true

	//	fmt.Println(user)                //mengambil semua data
	//	fmt.Println(user.ID, user.Email) //mengambil sebagian data
	//	fmt.Println(user2)

	//penulisan cara ke 2
	user := User{
		ID:        1,
		Fisrtname: "joko",
		Lastname:  "ardiyanto",
		Email:     "ganteng@yahoo.co.id",
		Isactive:  true,
	}

	fmt.Println(user)

	//penulisan 3
	person := User{1, "zayn", "malik", "zayn@gmail.com", true}

	fmt.Println("========version 2===========")
	fmt.Println(person)

}
