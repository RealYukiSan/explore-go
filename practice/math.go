package practice

import (
	"fmt"
	"math"
)

func Math() {
	fmt.Printf("Results : %d\n", gradingStudents([]uint8{99, 100, 91, 38, 37, 45, 78}))
}

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */
func gradingStudents(grades []uint8) []uint8 {
	// Write your code here
	for i, x := range grades {
		difference := uint8(math.Ceil(float64(x)/5) * 5)

		if x >= 38 && x <= 100 && difference-x < 3 {
			grades[i] = difference
		}
	}

	return grades
}
