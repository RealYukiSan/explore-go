package test

import (
	"fmt"
	"runtime"
)

func BufferedChannel() {
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 3)

	go func() {
		for {
			i := <-messages
			fmt.Println("receive data", i)

		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i

	}

	// kode dibawah berperan untuk membiarkan data terakhir agar diterima oleh goroutine
	// time.Sleep(5 * time.Second)
}
