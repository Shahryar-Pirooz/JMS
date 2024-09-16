package main

import (
	"fmt"
	"os"
)

var (
	Root string
)

func init() {
	if root, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		Root = root
	}
}

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
