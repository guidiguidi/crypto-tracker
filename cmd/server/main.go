package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/guidiguidi/crypto-tracker/internal/config"
	"github.com/guidiguidi/crypto-tracker/internal/repository"
	"github.com/guidiguidi/crypto-tracker/internal/api/handlers"
)

func main() {
	cfg := config.Load()
	repo, err := repository.New(cfg.DB.URL)
	if err != nil {
		log.Fatal("DB failed:", err)
	}

	r := gin.Default()
	r.GET("/health", handlers.Health)
	r.POST("/portfolio", handlers.CreatePortfolio(repo))
	r.GET("/portfolio/:user_id", handlers.GetPortfolio(repo)) // Новый!
	r.Run(":" + cfg.Server.Port)
}
