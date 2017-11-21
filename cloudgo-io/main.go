package main

import (
	"os"

	"github.com/James-Yip/service_computing_homework/cloudgo-io/service"
	flag "github.com/spf13/pflag"
)

const (
	DEFAULT_PORT string = "8080"
)

func main() {
	var port string
	// get the port inputted by user
	flag.StringVarP(&port, "port", "p", "", "cloudgo listening port")
	flag.Parse()
	// if user does not specify port
	if port == "" {
		// get current environment variable PORT
		envPort := os.Getenv("PORT")
		// if PORT is set
		if envPort != "" {
			// set the listening port to PORT
			port = envPort
		} else {
			// set the listening port to default value (8080)
			port = DEFAULT_PORT
		}
	}
	service.RunServer(port)
}
