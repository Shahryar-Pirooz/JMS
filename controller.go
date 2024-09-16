package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

var data []map[string]interface{}

func jsonReader(urlPath string, formatRequest string) error {
	fullPath := Root + urlPath + formatRequest
	file, err := os.Open(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return err
	}
	return nil
}

func GetRoute(ctx *gin.Context) {
	jsonReader(ctx.Request.URL.Path, "_get.json")
	ctx.JSON(200, data)
}

func PostRoute(ctx *gin.Context) {
	jsonReader(ctx.Request.URL.Path, "_post.json")
	ctx.String(200, "Hi")

}
