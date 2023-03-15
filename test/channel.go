package test

import (
	"fmt"
	"runtime"
)

func Channel() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)
	// fungsi ini dijalankan secara asynchronus karena ia merupakan sebuah fungsi yang dijalankan pada goroutine lain
	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("Hello %s", who)
		messages <- data
	}

	go sayHelloTo("john wick")
	go sayHelloTo("makise chan")
	go sayHelloTo("shinigami")

	// tetapi serah-terimanya dilakukan secara synchronus, maka dari itu kita tidak perlu membuat code blocking seperti fmt.Scan pada commit sebelumnya
	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)
}
