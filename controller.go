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

// jsonReader reads JSON data from a file and decodes it into the response variable
// urlPath: the path of the URL being accessed
// formatRequest: the format of the request (e.g., "_get.json")
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

// jsonWriter writes the request data to a JSON file
// urlPath: the path of the URL being accessed
// formatRequest: the format of the request (e.g., "_post.json")
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

// GetRoute handles GET requests
// It reads data from a JSON file and sends it as a response
func GetRoute(ctx *gin.Context) {
	if err := jsonReader(ctx.Request.URL.Path, "_get.json"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// PostRoute handles POST requests
// It reads the request body, writes it to a JSON file, and sends a confirmation response
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
