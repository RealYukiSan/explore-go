package main

import (
	"explore-go/practice"
	"fmt"
)

func main() {
	go practice.RestFulAPIServer()
	users, err := practice.HttpClient()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range users {
		fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	}

	fmt.Scanln()
}
