package service

import (
	"context"

	"github.com/aveplen-bach/config-service/internal/model"
)

func (c *ConfigurationService) GetFacerecConfig(ctx context.Context) (model.FacerecConfig, error) {
	return c.repo.GetFaceRecognitionConfig(ctx)
}

func (c *ConfigurationService) UpdateFacerecConfig(ctx context.Context, config model.FacerecConfig) error {
	return c.repo.UpdateFaceRecognitionConfig(ctx, config)
}
