package repo

import (
	"context"
	"github.com/aveplen-bach/config-service/internal/model"
)

func (r *Repository) GetAuthenticationDatabaseConfig(ctx context.Context) (model.AuthenticationDatabaseConfig, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var config model.AuthenticationDatabaseConfig
	if err := r.db.WithContext(ctx).Order("created_at").Last(&config).Error; err != nil {
		return model.AuthenticationDatabaseConfig{}, err
	}

	return config, nil
}

func (r *Repository) UpdateAuthenticationDatabaseConfig(ctx context.Context, config model.AuthenticationDatabaseConfig) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := r.db.WithContext(ctx).Save(&config).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetAuthenticationRegistration(ctx context.Context) (model.AuthenticationRegistration, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var config model.AuthenticationRegistration
	if err := r.db.WithContext(ctx).Order("created_at").Last(&config).Error; err != nil {
		return model.AuthenticationRegistration{}, err
	}

	return config, nil
}

func (r *Repository) UpdateAuthenticationRegistration(ctx context.Context, registration model.AuthenticationRegistration) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := r.db.WithContext(ctx).Save(&registration).Error; err != nil {
		return err
	}

	return nil
}
