package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

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
