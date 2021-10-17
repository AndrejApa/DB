package main

import (
	"net/http"
	"time"
)

type User struct {
	Id       uint      `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	BirthDay time.Time `json:"birthDay"`
}

func getPersonByID(w http.ResponseWriter, r *http.Request) {

}
func updatePersonByID(w http.ResponseWriter, r *http.Request) {

}
func deletePersonByID(w http.ResponseWriter, r *http.Request) {

}
func main() {

}
