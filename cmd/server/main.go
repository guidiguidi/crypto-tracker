package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health", healthHandler)
	r.GET("/prices", pricesHandler)
	r.Run(":8080")
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func pricesHandler(c *gin.Context) {
	coins := c.Query("coins")
	if coins == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "coins query parameter is required"})
		return
	}
	coinList := splitCoins(coins)
	prices, err := fetchCoinGecko(coinList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch prices"})
		return
	}
	c.JSON(http.StatusOK, prices)
	
}
func splitCoins(coins string) []string {
	var coinList []string
	for _, coin := range strings.Split(coins, ",") {
		coinList = append(coinList, coin)
	}
	return coinList
}

func fetchCoinGecko(coins []string) (map[string]interface{}, error) {
    ids := strings.Join(coins, ",")
    url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", ids)

    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("HTTP request failed: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("read response failed: %w", err)
    }

    var prices map[string]interface{}
    if err := json.Unmarshal(body, &prices); err != nil {
        return nil, fmt.Errorf("JSON unmarshal failed: %w", err)
    }

    return prices, nil
}