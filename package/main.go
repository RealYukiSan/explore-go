package main

import (
	"explore-go/test"
	"fmt"
)

func main() {
	coba := test.Oke{Exposed: "dapat diakses lhoo"}
	fmt.Println(&coba.Exposed)
}
