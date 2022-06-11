package repo

import (
	"github.com/aveplen-bach/config-service/internal/model"
	"golang.org/x/net/context"
)

func (r *Repository) GetFaceRecognitionConfig(ctx context.Context) (model.FacerecConfig, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var config model.FacerecConfig
	if err := r.db.WithContext(ctx).Order("created_at").Last(&config).Error; err != nil {
		return model.FacerecConfig{}, err
	}

	return config, nil
}

func (r *Repository) UpdateFaceRecognitionConfig(ctx context.Context, config model.FacerecConfig) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := r.db.WithContext(ctx).Save(&config).Error; err != nil {
		return err
	}

	return nil
}
