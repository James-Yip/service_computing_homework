package service

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

type User struct {
	Username string
	Password string
	Phone    string
	Email    string
}

func signUp(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		// send signup.html detailto client
		t, _ := template.ParseFiles("assets/html/signup.html")
		t.Execute(rw, nil)
	} else {
		req.ParseForm()
		// convert the request form (map) into struct User
		user := new(User)
		decoder := schema.NewDecoder()
		if err := decoder.Decode(user, req.Form); err != nil {
			log.Fatal(err)
		}
		// send a HTML that contains the user information
		t, err := template.ParseFiles("templates/detail.gtpl")
		if err != nil {
			log.Fatal(err)
		}
		if err = t.Execute(rw, *user); err != nil {
			log.Fatal(err)
		}
	}
}
