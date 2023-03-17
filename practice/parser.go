package practice

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func ImplementParser() {
	parseJson()
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

// struct untuk menampung hasil decode json nya
type User struct {
	Fullname string `json:"Name"`
	Age      int
}

func parseJson() {
	jsonString := `{"Name": "Yare-yare", "Age": 27}`
	arrayJsonString := `[
		{"Name": "Yare-yare", "Age": 27},
		{"Name": "Ishida", "Age": 21},
		{"Name": "Daru", "Age": 17}
	]`
	jsonData := []byte(jsonString)
	var data User
	var data1 map[string]interface{}
	var data2 interface{}
	var data3 []User

	// decode section
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Converted to struct")
	fmt.Println("Fullname :", data.Fullname)
	fmt.Println("Age :", data.Age)

	e := json.Unmarshal(jsonData, &data1)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Println("Converted to map")
	fmt.Println("Fullname :", data1["Name"])
	fmt.Println("Age :", data1["Age"])

	er := json.Unmarshal(jsonData, &data2)
	if er != nil {
		fmt.Println(er.Error())
		return
	}

	decodedData := data2.(map[string]interface{})
	fmt.Println("Converted to interface")
	fmt.Println("Fullname :", decodedData["Name"])
	fmt.Println("Age :", decodedData["Age"])

	fmt.Println("Convert array json to slice")
	json.Unmarshal([]byte(arrayJsonString), &data3)

	for _, user := range data3 {
		fmt.Println(user.Fullname)
	}

	// encode section
	var object = []User{{"L", 12}, {"Yagami", 23}}
	var jesonData, emror = json.Marshal(object)
	if emror != nil {
		fmt.Println(emror.Error())
		return
	}

	var jesonString = string(jesonData)
	fmt.Println(jesonString)
}
