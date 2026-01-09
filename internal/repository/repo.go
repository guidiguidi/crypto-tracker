package repository

import (
	"github.com/guidiguidi/crypto-tracker/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate
	db.AutoMigrate(&models.Portfolio{})
	return &Repository{db: db}, nil
}

func (r *Repository) CreatePortfolio(p *models.Portfolio) error {
	return r.db.Create(p).Error
}

func (r *Repository) GetPortfolio(userID uint) ([]models.Portfolio, error) {
	var portfolios []models.Portfolio
	err := r.db.Where("user_id = ?", userID).Find(&portfolios).Error
	return portfolios, err
}
