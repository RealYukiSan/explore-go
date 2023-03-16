package test

import "fmt"

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

	// multi-dimension array
	matrix := [2][3]int{[3]int{3, 2, 1}, [3]int{1, 2, 3}}

	fmt.Println("matrix :", matrix)
}
