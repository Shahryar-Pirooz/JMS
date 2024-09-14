package main

import (
	"errors"
	"io/fs"
	"path/filepath"
	"strings"
)

var jsonAddress []string

func walkFunc(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
		jsonAddress = append(jsonAddress, path)
	}
	return nil
}

func GetJson() ([]string, error) {
	err := filepath.Walk(Root, walkFunc)
	if err != nil {
		panic(err)
	}
	if len(jsonAddress) == 0 {
		return nil, errors.New("can not find any json")
	}
	return jsonAddress, nil
}
