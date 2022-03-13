// membuat Method Pointer Receiver
package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func (s *Student) graduate() { // diberi * pada student
	s.Name = s.Name + " S.Kom"
}

//ketika ingin buat func di dalamnya melakukan kalkulasi yg mempengaruhi si soffer (s) di method func
//maka student yang struct harus dikasi tanda *

func main() {
	student := Student{1, "joko ardi", 3.50}

	fmt.Println(student.Name)

	student.graduate()
	fmt.Println(student.Name)
}
