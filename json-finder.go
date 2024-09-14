package main

import (
	"errors"
	"io/fs"
	"path/filepath"
	"strings"
)

type JSON struct {
	POST []string
	GET  []string
}

var (
	jsonAddress   JSON
	folderAddress []string
)

func walkFunc(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		folderAddress = append(folderAddress, path)
	}
	if !info.IsDir() {
		if strings.HasSuffix(info.Name(), "post.json") {
			jsonAddress.POST = append(jsonAddress.POST, path)
		}
		if strings.HasSuffix(info.Name(), "get.json") {
			jsonAddress.GET = append(jsonAddress.GET, path)

		}
	}
	return nil
}

func Walker() (JSON, []string, error) {
	err := filepath.Walk(Root, walkFunc)
	if err != nil {
		panic(err)
	}
	if len(jsonAddress.POST) == 0 && len(jsonAddress.GET) == 0 {
		return JSON{}, nil, errors.New("can not find any json")
	}
	return jsonAddress, folderAddress, nil
}
