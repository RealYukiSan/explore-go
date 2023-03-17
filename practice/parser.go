package practice

import (
	"fmt"
	"net/url"
)

func ImplementParser() {
	urlParsing()
}

func urlParsing() {
	var urlString = "http://kalipare.com:80/hello?name=hon wick&age=27"
	var u, err = url.Parse(urlString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)

	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)

	name := u.Query()["name"][0]
	age := u.Query()["age"][0]
	fmt.Printf("name: %s, age: %s\n", name, age)

}
