package fundamental

import "fmt"

func ImplementFunction() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	// example of closure
	increment := incremental(numbers...)

	increment()
	increment()
	increment()
	increment()
	increment()
	fmt.Println(increment())
}

// variadic function cuz have rest parameter, perlu di ingat: pastikan jika ingin menggunakan rest parameter, taruh di urutan terakhir
func incremental(numbers ...int) func() int {
	i := 0

	// loop tidak akan dijalankan jika numbersnya kosong, tetapi ternyata tidak error bjir kalo waktu manggil fungsi ini ga masukin parameternya, berbeda dengan typescript ygy :v
	for _, number := range numbers {
		fmt.Println("number \t:", number)
	}

	return func() int {
		i++
		return i
	}
}
