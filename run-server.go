package main

import (
	"github.com/gin-gonic/gin"
)

// RunServer initializes and starts the Gin server
// It sets up the routes and runs the server on the specified port
func RunServer(port string) {
	gin.SetMode(gin.ReleaseMode)
	json := Walker()
	r := gin.Default()
	RouteCreator(json, r)
	r.Run(port)
}
