package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sapConnector/models"
)

func initializeRoutes() {
	// V1 router definition
	v1 := r.Group("/v1")
	{
		v1.GET("/status", apiStatusCheck) // Test path

		// "v1/products" routes
		productRoutes := v1.Group("/products")
		{
			productRoutes.GET("/", getAllProducts)
			productRoutes.GET("/:product_id", getProduct)
		}
	}
}

// Respond with 200 code and ok JSON
func apiStatusCheck(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"api status": "OK",
		})
}

// Respond with complete product list
func getAllProducts(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"products": models.GetAllProducts(),
		})
}

// Respond with the selected product by id
func getProduct(c *gin.Context) {
	if productId, err := models.GetProductById(c.Param("product_id")); err == nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"param": productId,
			})
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
}
