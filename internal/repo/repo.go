package repo

import (
	"sync"

	"github.com/aveplen-bach/config-service/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
	mu *sync.Mutex
}

func NewRepository() (*Repository, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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
