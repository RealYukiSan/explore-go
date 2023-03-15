package practice

import (
	"fmt"
	"math"
)

func Math() {
	gradingStudents()
}

func gradingStudents() {
	for _, x := range []int{36, 38, 100, 98, 61} {
		difference := int(math.Ceil(float64(x)/5) * 5)

		if x >= 38 && x <= 100 && difference-x < 3 {
			fmt.Printf("Output : %d\n", difference)
		} else {
			fmt.Println("The number cannot ceiled :v")
		}
	}
}
