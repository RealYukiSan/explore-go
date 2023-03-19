package main

import (
	"explore-go/fundamental"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	fundamental.ImplementMutex()

	// go practice.RestFulAPIServer()
	// practice.HttpClient()
	// fmt.Scanln()
}
