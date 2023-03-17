package practice

import (
	"fmt"
	"math"
)

func Math() {
	// fmt.Printf("Results : %d\n", gradingStudents([]uint8{99, 100, 91, 38, 37, 45, 78}))
	plusMinus([]int32{1, 2, 0, -4, 5})
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

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
func plusMinus(arr []int32) {
	// Write your code here
	positive := 0
	negative := 0
	zero := 0
	for _, value := range arr {
		if value < 0 {
			negative++
		} else if value > 0 {
			positive++
		} else {
			zero++
		}
	}

	fmt.Printf("%f\n", float32(positive)/float32(len(arr)))
	fmt.Printf("%f\n", float32(negative)/float32(len(arr)))
	fmt.Printf("%f\n", float32(zero)/float32(len(arr)))
}
