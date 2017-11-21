package service

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

// RunServer initialize a server and run it at certain port
func RunServer(port string) {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	// initial Negroni with default middlewares
	n := negroni.Classic()
	// create router
	router := mux.NewRouter()
	// get current web server resource path
	webRoot := getWebRoot()
	// register route
	router.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/signup", signUp)
	router.HandleFunc("/unknown", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(501)
	})
	// provide static file access service
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir(webRoot+"/assets/"))))
	// use certain router
	n.UseHandler(router)
	n.Run(":" + port)
}

func getWebRoot() string {
	// get current environment variable "WEBROOT"
	webRoot := os.Getenv("WEBROOT")
	if webRoot == "" {
		// get current working directory
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}
	return webRoot
}
