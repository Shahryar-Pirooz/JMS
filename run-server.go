package main

import (
	"github.com/gin-gonic/gin"
)

func RunServer(port string) {
	gin.SetMode(gin.ReleaseMode)
	json := Walker()
	r := gin.Default()
	RouteCreator(json, r)
	r.Run(port)
}
