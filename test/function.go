package test

import "fmt"

func ImplementFunction() {
	// example of closure
	increment := incremental()
	increment()
	increment()
	increment()
	increment()
	increment()
	fmt.Println(increment())
}

func incremental() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
