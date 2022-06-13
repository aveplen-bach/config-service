package repo

import (
	"fmt"

	"github.com/aveplen-bach/config-service/internal/model"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

func (r *Repository) GetFaceRecognitionConfig(ctx context.Context) (model.FacerecConfig, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	logrus.Info("getting last facerec config from db")
	var config model.FacerecConfig
	if err := r.db.WithContext(ctx).Order("created_at").Last(&config).Error; err != nil {
		logrus.Warnf("could not get last facerec config from db: %w", err)
		return model.FacerecConfig{}, fmt.Errorf("could not get last facerec config from db: %w", err)
	}

	return config, nil
}

func (r *Repository) UpdateFaceRecognitionConfig(ctx context.Context, config model.FacerecConfig) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	logrus.Info("saving new facerec config to db")
	if err := r.db.WithContext(ctx).Save(&config).Error; err != nil {
		logrus.Warnf("could not save new facerec config to db: %w", err)
		return fmt.Errorf("could not save new facerec config to db: %w", err)
	}

	return nil
}
