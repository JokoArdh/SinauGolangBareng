package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func main() {

	student := Student{1, "joko ardiyanto", 3.50}

	fmt.Println(student.Name)

	//kasus dalam sebuah nama , jiki sudah lulus maka menyandang gelar s.kom
	//dari hal itu kita akan mempersing nilai dengan menggunakan pinter

	graduate(&student) //function graduate kita parsing & untuk alamat memori
	fmt.Println(student.Name)

}

func graduate(student *Student) { // kita parsing * untuk mengambil nilai
	student.Name = student.Name + " S.Kom"
}
