package fundamental

import (
	"fmt"
)

func Channel() {
	var messages = make(chan string)
	for _, each := range []string{"lorea", "kiana", "kurisu"} {
		// fungsi ini dijalankan secara asynchronus karena ia merupakan sebuah fungsi yang dijalankan pada goroutine lain
		go func(who string) {
			var data = fmt.Sprintf("Hello %s", who)
			messages <- data
		}(each)
	}

	// tetapi serah-terimanya dilakukan secara synchronus, maka dari itu kita tidak perlu membuat code blocking seperti fmt.Scan pada commit sebelumnya
	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}

func printMessage(message chan string) {
	fmt.Println(<-message)
}
