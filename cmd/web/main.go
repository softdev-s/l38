package main

import (
	"fmt"
	"github.com/softdev-s/l38/pkg/handlers"
	"net/http"
)

const portNumber = ":2345"

// main is the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
