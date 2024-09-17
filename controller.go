package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Data map[string]interface{}

var response []Data
var request Data

func jsonReader(urlPath string, formatRequest string) error {
	fullPath := Root + urlPath + formatRequest
	file, err := os.Open(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&response)
	if err != nil {
		return err
	}
	return nil
}

func jsonWriter(urlPath string, formatRequest string) error {
	fullPath := Root + urlPath + formatRequest

	var existingData []Data
	existingContent, err := os.ReadFile(fullPath)
	if err == nil {
		if err := json.Unmarshal(existingContent, &existingData); err != nil {
			return err
		}
	}
	existingData = append(existingData, request)

	updatedContent, err := json.MarshalIndent(existingData, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fullPath, updatedContent, 0644)
}

func GetRoute(ctx *gin.Context) {
	if err := jsonReader(ctx.Request.URL.Path, "_get.json"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func PostRoute(ctx *gin.Context) {
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	if err := jsonWriter(ctx.Request.URL.Path, "_post.json"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "added",
		"data":    request,
	})

}
