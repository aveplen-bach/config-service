package repo

import (
	"github.com/aveplen-bach/config-service/internal/model"
	"golang.org/x/net/context"
)

func (r *Repository) GetFaceRecognitionConfig(ctx context.Context) (model.FaceRecognitionConfig, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var config model.FaceRecognitionConfig
	if err := r.db.WithContext(ctx).Order("created_at").Last(&config).Error; err != nil {
		return model.FaceRecognitionConfig{}, err
	}

	return config, nil
}

func (r *Repository) UpdateFaceRecognitionConfig(ctx context.Context, config model.FaceRecognitionConfig) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := r.db.WithContext(ctx).Save(&config).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetFaceRecognitionRegistration(ctx context.Context) (model.FaceRecognitionRegistration, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var config model.FaceRecognitionRegistration
	if err := r.db.WithContext(ctx).Order("created_at").Last(&config).Error; err != nil {
		return model.FaceRecognitionRegistration{}, err
	}

	return config, nil
}

func (r *Repository) UpdateFaceRecognitionRegistration(ctx context.Context, registration model.FaceRecognitionRegistration) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := r.db.WithContext(ctx).Save(&registration).Error; err != nil {
		return err
	}

	return nil
}
