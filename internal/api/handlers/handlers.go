package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guidiguidi/crypto-tracker/internal/models"
	"github.com/guidiguidi/crypto-tracker/internal/repository"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"}) // StatusOK = 200
}

func CreatePortfolio(repo *repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p models.Portfolio
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := repo.CreatePortfolio(&p); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, p)
	}
}

func GetPortfolio(repo *repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr := c.Param("user_id")
		userID, _ := strconv.Atoi(userIDStr)
		portfolios, err := repo.GetPortfolio(uint(userID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, portfolios)
	}
}
