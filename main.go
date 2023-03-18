package main

import (
	"explore-go/practice"
	"fmt"
)

func main() {
	go practice.RestFulAPIServer()
	practice.HttpClient()
	fmt.Scanln()
}
