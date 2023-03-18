package fundamental

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func Test() {
	x := T{1, "test"}
	p := reflect.ValueOf(&x).Elem()
	typeOfX := p.Type()
	for i := 0; i < p.NumField(); i++ {
		fmt.Printf("%d: %s %s = %v\n", i, typeOfX.Field(i).Name, p.Field(i).Type(), p.Field(i).Interface())
	}

	p.Field(1).SetString("modified!")
	fmt.Printf("x is now: %v\n", x)
}
