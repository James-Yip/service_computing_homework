package service

import (
	"html/template"
	"net/http"
)

func homeHandler(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("assets/html/index.html")
	t.Execute(rw, nil)
}
