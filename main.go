package main

import (
	"fmt"
	"os"
)

var (
	Root string
)

// init initializes the Root variable with the current working directory
func init() {
	if root, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		Root = root
	}
}

// main is the entry point of the application
// It determines the port to run the server on and starts the server
func main() {
	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = ":3000"
	}
	fmt.Println("JMS is run on http://localhost" + port)
	RunServer(port)
}
