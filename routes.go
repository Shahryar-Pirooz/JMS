package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

// RouteCreator sets up the GET and POST routes for the Gin engine
// It prints the endpoints and associates them with the appropriate handler functions
func RouteCreator(json JSON, g *gin.Engine) {
	fmt.Println("Your GET endpoints :")
	for _, dir := range json.GET {
		d := strings.TrimPrefix(dir, Root)
		d = strings.TrimSuffix(d, "_get.json")
		color.Green(d)
		g.GET(d, GetRoute)
	}
	fmt.Println("Your POST endpoints :")
	for _, dir := range json.POST {
		d := strings.TrimPrefix(dir, Root)
		d = strings.TrimSuffix(d, "_post.json")
		color.Blue(d)
		g.POST(d, PostRoute)
	}
}
