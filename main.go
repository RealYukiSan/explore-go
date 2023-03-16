package main

import (
	"explore-go/test"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	test.TimeoutChannel()
}
