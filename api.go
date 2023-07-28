package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name  string `json: "name"`
	Age   int    `json: "age"`
	Hobby string `json: "hobby"`
}

var users = []User{{Name: "user1", Age: 35, Hobby: "reading"}}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)

}

func AddUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	users = append(users, user)

	json.NewEncoder(w).Encode(user)
}
func main() {

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/users", GetUsers)
	r.HandleFunc("/user", AddUser).Methods("POST")
	// r.HandleFunc("")
	fmt.Println("Server started")

	log.Fatal(http.ListenAndServe(":10000", r))

}
