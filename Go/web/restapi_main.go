// restapi_main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID   int
	Name string
}

var users = []User{
	{ID: 1, Name: "张三"},
	{ID: 2, Name: "李四"},
	{ID: 2, Name: "王五"},
}

func main() {
	http.HandleFunc("/users", handleUsers)
	http.ListenAndServe(":8080", nil)
	//fmt.Println("Hello World!")
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "{\"message\": \""+err.Error()+"\"}")
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(users)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "{\"message\": \"not found\"}")
	}
}
