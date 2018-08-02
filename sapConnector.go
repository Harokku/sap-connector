package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Router definition
	r := gin.Default()

	// Default status check route
	r.GET("/status", apiStatusCheck)

	// V1 router definition
	v1 := r.Group("/v1")
	{
		v1.GET("/status", apiStatusCheck)
	}

	// Run the router
	r.Run()

}
func apiStatusCheck(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"Api status": "OK",
		})
}
