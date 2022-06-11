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
	db, err := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(model.FacerecConfig{})

	return &Repository{
		db: db,
		mu: &sync.Mutex{},
	}, nil
}
