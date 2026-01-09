package handlers

import (
	"github.com/guidiguidi/crypto-tracker/internal/models"
	"github.com/guidiguidi/crypto-tracker/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createPortfolioHandler(repo *repository.Repository) gin.HandlerFunc {
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
