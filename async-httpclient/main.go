package main

import (
	"fmt"

	"github.com/James-Yip/service_computing_homework/async-httpclient/scenario"
)

func main() {
	fmt.Println("scenario1: synchronous way")
	scenario.Synchronous()
	fmt.Println("scenario1: asynchronous way")
	scenario.Asynchronous()
}
