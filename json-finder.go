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

func isContains(arr []string, item string) bool {
	for _, str := range arr {
		if str == item {
			return true
		}
	}
	return false
}

func walkFunc(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
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

func Walker() JSON {
	err := filepath.Walk(Root, walkFunc)
	if err != nil {
		panic(err)
	}
	return jsonAddress
}
