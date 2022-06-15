package repo

import (
	"fmt"
	"sync"

	"github.com/aveplen-bach/config-service/internal/config"
	"github.com/aveplen-bach/config-service/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
	mu *sync.Mutex
}

func NewRepository(cfg config.Config) (*Repository, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DatabaseConfig.Host,
		cfg.DatabaseConfig.User,
		cfg.DatabaseConfig.Password,
		cfg.DatabaseConfig.Name,
		cfg.DatabaseConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(model.FacerecConfig{})

	var count int64
	db.Model(&model.FacerecConfig{}).Count(&count)
	if count == 0 {
		db.Save(&model.FacerecConfig{Threshold: 0.6})
	}

	return &Repository{
		db: db,
		mu: &sync.Mutex{},
	}, nil
}
