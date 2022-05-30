package repo

import (
	"github.com/aveplen-bach/config-service/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
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

	db.AutoMigrate(model.AuthenticationRegistration{})
	db.AutoMigrate(model.AuthenticationDatabaseConfig{})

	db.AutoMigrate(model.FaceRecognitionConfig{})
	db.AutoMigrate(model.FaceRecognitionRegistration{})

	db.AutoMigrate(model.S3GatewayConfig{})
	db.AutoMigrate(model.S3GatewayRegistration{})

	return &Repository{
		db: db,
		mu: &sync.Mutex{},
	}, nil
}
