package main

import "fmt"

func main() {

	fmt.Println("Main function")

	var studentTemp Student
	studentTemp = Student{}
	fmt.Println("student before init ", studentTemp)

	studentTemp.Address = Address{"Nalanda Colony", "Thana", "Patna", "Bihar"}
	studentTemp.Id = "12334"
	studentTemp.Name = "Rupesh"
	studentTemp.Email = "Rupesh@gmail.com"
	fmt.Println("student before init ", studentTemp)
}

type Student struct {
	Id      string
	Name    string
	Email   string
	Address Address
	Course  *Course
}

type Address struct {
	Muhalla  string
	Thana    string
	District string
	State    string
}
type Course struct {
	Name string
}
