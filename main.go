package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/health", healthHandler)
	r.GET("/prices", pricesHandler)
	r.Run(":8080")
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func pricesHandler(c *gin.Context) {
	coins := c.Query("coins")
	_ = coins 

	c.JSON(http.StatusOK, gin.H{"prices": map[string]float64{}})

}
