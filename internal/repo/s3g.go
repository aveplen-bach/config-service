package repo

import (
	"context"
	"github.com/aveplen-bach/config-service/internal/model"
)

func (r *Repository) GetS3GatewayConfig(ctx context.Context) (model.S3GatewayConfig, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var config model.S3GatewayConfig
	if err := r.db.WithContext(ctx).Order("created_at").Last(&config).Error; err != nil {
		return model.S3GatewayConfig{}, err
	}

	return config, nil
}

func (r *Repository) UpdateS3GatewayConfigConfig(ctx context.Context, config model.S3GatewayConfig) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := r.db.WithContext(ctx).Save(&config).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetS3GatewayRegistration(ctx context.Context) (model.S3GatewayRegistration, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var config model.S3GatewayRegistration
	if err := r.db.WithContext(ctx).Order("created_at").Last(&config).Error; err != nil {
		return model.S3GatewayRegistration{}, err
	}

	return config, nil
}

func (r *Repository) UpdateS3GatewayRegistration(ctx context.Context, registration model.S3GatewayRegistration) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := r.db.WithContext(ctx).Save(&registration).Error; err != nil {
		return err
	}

	return nil
}
