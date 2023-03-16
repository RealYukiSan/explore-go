package test

import (
	"fmt"
	"reflect"
)

func ImplementArray() {
	var names [4]string
	names[0] = "kurisu"
	names[1] = "de"
	names[2] = "rune"
	names[3] = "desu"

	// more scallable rather than hard code with each index :v
	for i, item := range names {
		fmt.Print(item, " ")
		if i == 3 {
			fmt.Print("\n")
		}
	}

	// init type array and assign it with default values
	fruits := [4]string{"apel", "semangka", "melon", "banana"}
	fmt.Println(fruits)

	// set length to automatically adjust with the length of items!
	numbers := [...]int{1, 2, 3, 4}

	fmt.Println("data array \t:", numbers)
	fmt.Println("type \t:", reflect.TypeOf(numbers).Kind())

	// multi-dimension array
	matrix := [2][3]int{{3, 2, 1}, {1, 2, 3}}

	fmt.Println("matrix :", matrix)
}

func ImplementSlice() {
	var fruits = make([]string, 2)
	fruits[0] = "khuldi"
	fruits[1] = "sirsak"
	fmt.Println("type \t:", reflect.TypeOf(fruits).Kind())

	fruits = []string{"apple", "grape", "banana"}
	aFruits := fruits[:3]
	// index terakhir pada slice ini menandakan jumlah kapasitas yang ditentukan
	bFruits := fruits[2:3:3]
	cFruits := fruits[2:]
	dFruits := aFruits[:]

	cFruits[0] = "changed"
	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits, len(bFruits), cap(bFruits))
	fmt.Println(cFruits)
	fmt.Println(dFruits)
}
