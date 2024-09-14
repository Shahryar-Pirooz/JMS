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
	// RunServer()
	json, err := GetJson()
	if err != nil {
		panic(err)
	}
	fmt.Println(json)
}
