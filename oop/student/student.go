package student

import (
	"fmt"
)

type student struct {
	id    string
	name  string
	email string
}

func New(id string, name string, email string) student {
	stu := student{id, name, email}
	return stu
}

func (stu student) ToString() {
	fmt.Println("id:" + stu.id + " name:" + stu.name + " email:" + stu.email)
}
