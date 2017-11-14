package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// NewServer configures and returns a Server.
func RunServer(port string) {
	// initial Negroni with default middlewares
	n := negroni.Classic()
	// create router
	router := mux.NewRouter()
	// register route
	router.HandleFunc("/", showPrompt)
	router.HandleFunc("/{username}", sayHello)
	// use certain router
	n.UseHandler(router)
	n.Run(":" + port)
}

func showPrompt(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "This is cloudgo.")
	fmt.Fprintln(rw, "Please input your name in the url.")
	fmt.Fprintln(rw, "example: localhost:8080/James")
}

func sayHello(rw http.ResponseWriter, req *http.Request) {
	// get the variables from request url
	vars := mux.Vars(req)
	// extract the user name
	username := vars["username"]
	// write data into ResponseWriter
	fmt.Fprintf(rw, "Hello %s!\n", username)
}
