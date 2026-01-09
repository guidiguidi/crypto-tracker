package models

import "gorm.io/gorm"

type Portfolio struct {
	gorm.Model
	UserID   uint    `json:"user_id" gorm:"index"`
	Coin     string  `json:"coin"`
	Amount   float64 `json:"amount"`
	AvgPrice float64 `json:"avg_price"`
}
