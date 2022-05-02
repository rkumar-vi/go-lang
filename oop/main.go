package main

import "oop/student"

func main() {

	// print Student Data
	student_1 := student.New("123", "rupesh kumar", "rupesh.kumar@encora.com")

	student_1.ToString()

}

// ouptput
//rupeshkumar@ENCOPDBANLT0254 oop % pwd
///Users/rupeshkumar/go-lang/oop
//rupeshkumar@ENCOPDBANLT0254 oop % go run *.go
//id:123 name:rupesh kumar email:rupesh.kumar@encora.com
//rupeshkumar@ENCOPDBANLT0254 oop %
