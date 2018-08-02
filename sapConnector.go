package main

import (
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {
	// Router definition
	r = gin.Default()

	// Default status check route
	r.GET("/status", apiStatusCheck)

	// Initialize routes from routes.go
	initializeRoutes()

	// Run the router
	r.Run()

}
