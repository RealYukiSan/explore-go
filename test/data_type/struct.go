package test

import "fmt"

type student struct {
	name  string
	grade int
}

func ImplementStruct() {
	var s1 student
	s1.grade = 1
	s1.name = "John Titor"

	var s2 = student{name: "kyoma"}

	var s3 = student{"kurisu", 3}

	fmt.Println("student 1 :", s1.name)
	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)

	fmt.Println("zero value of s2 :", s2.grade)
}
