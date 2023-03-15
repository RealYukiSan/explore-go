package main

import (
	"explore-go/test"
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	coba := test.Oke{Exposed: "dapat diakses lhoo"}
	fmt.Println(&coba.Exposed)

	test.Test()
	go test.Print(5, "aloo")
	test.Print(5, "no goroutine")

	var test string
	fmt.Scanln(&test)
	fmt.Println(test)
}
