package main

import (
	"explore-go/fundamental"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	fundamental.ImplementWaitGroup()

	// go practice.RestFulAPIServer()
	// practice.HttpClient()
	// fmt.Scanln()
}
