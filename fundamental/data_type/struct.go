package fundamental

import "fmt"

type person struct {
	name string
	age  uint8
}

type student struct {
	person
	grade int
}

func ImplementStruct() {
	var s1 = student{}
	s1.grade = 1
	s1.name = "John Titor"
	s1.age = 23

	fmt.Println("name :", s1.name)
	fmt.Println("age :", s1.age)
	fmt.Println("grade :", s1.grade)
	// ini berguna jika parent dan child memiliki property yang sama, assign valuenya perlu dilakukan secara eksplisit!
	fmt.Println("another way to access child prop / sub struct :", s1.person.name)

	var p2 = person{"kami da", 18}
	var s2 = student{p2, 2}
	fmt.Println("it's also works :", s2)

	// anonymous struct!
	var s3 = struct {
		person
		grade int
	}{person{"third person", 19}, 3}

	fmt.Println("this is anonymous struct, useful if the struct only used once :", s3)

	// kombinasi antara slice dan struct, note: this combination also works with anonymous type, like the example above
	var allStudents = []student{
		s1, s2, s3,
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}
}
