package practice

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (_ *query) Hello() string { return "Hello word" }

func GQLServer() {
	s := `
		type Query {
			hello: String!
		}
	`

	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type student struct {
	ID    string
	Name  string
	Grade int
}

var data = []student{
	{
		ID:    "B001",
		Name:  "Testie",
		Grade: 1,
	},
	{
		ID:    "A002",
		Name:  "Kurisu",
		Grade: 3,
	},
	{
		ID:    "A003",
		Name:  "Kyoma",
		Grade: 2,
	},
}

const BASEURL = "http://localhost:8080"

func HttpClient() ([]student, error) {
	client := &http.Client{}
	var data []student
	var err error

	request, err := http.NewRequest("GET", BASEURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func RestFulAPIServer() {
	http.HandleFunc("/user", user)
	http.HandleFunc("/users", users)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		id := r.FormValue("id")
		for _, user := range data {
			if user.ID == id {
				var encoded, err = json.Marshal(data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(encoded)
				return
			}
		}

		http.Error(w, "User Not Found", http.StatusBadRequest)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}
func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		encoded, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(encoded)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
