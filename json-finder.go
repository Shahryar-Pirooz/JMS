package main

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type JSON struct {
	POST []string
	GET  []string
}

var (
	jsonAddress JSON
)

// isContains checks if a string is present in a slice of strings
func isContains(arr []string, item string) bool {
	for _, str := range arr {
		if str == item {
			return true
		}
	}
	return false
}

// walkFunc is used with filepath.Walk to process each file and directory
// It identifies JSON files for GET and POST requests and adds them to jsonAddress
func walkFunc(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
		// Normalize path to use forward slashes
		path = filepath.ToSlash(path)

		if strings.HasSuffix(info.Name(), "_post.json") {
			if !isContains(jsonAddress.POST, path) {
				jsonAddress.POST = append(jsonAddress.POST, path)
			}
		}
		if strings.HasSuffix(info.Name(), "_get.json") {
			if !isContains(jsonAddress.GET, path) {
				jsonAddress.GET = append(jsonAddress.GET, path)
			}
		}
	}
	return nil
}

// Walker traverses the directory structure starting from Root
// It collects paths to GET and POST JSON files and returns them in a JSON struct
func Walker() JSON {
	err := filepath.Walk(Root, walkFunc)
	if err != nil {
		panic(err)
	}
	return jsonAddress
}
