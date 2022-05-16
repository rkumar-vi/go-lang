package main

import "fmt"

func main() {

	fmt.Println("Main function")

	var studentTemp Student
	studentTemp = Student{"123", "Rupesh", "rupesh.kumar@gmail.com"}

	var objectRef Object
	objectRef = &studentTemp
	objectRef.ToString()
}

type Student struct {
	Id    string
	Name  string
	Email string
}

func (studentRef *Student) ToString() {
	fmt.Println("Id:" + studentRef.Id + " Name:" + studentRef.Name + " Email:" + studentRef.Email)

}

type Object interface {
	ToString()
}
