package service

import (
	"context"
	"github.com/aveplen-bach/config-service/internal/model"
)

func (c *ConfigurationService) GetFaceRecognitionConfig(ctx context.Context) (model.FaceRecognitionConfig, error) {
	return c.repo.GetFaceRecognitionConfig(ctx)
}

func (c *ConfigurationService) UpdateFaceRecognitionConfig(ctx context.Context, config model.FaceRecognitionConfig) error {
	return c.repo.UpdateFaceRecognitionConfig(ctx, config)
}

func (c *ConfigurationService) GetFaceRecognitionRegistration(ctx context.Context) (model.FaceRecognitionRegistration, error) {
	return c.repo.GetFaceRecognitionRegistration(ctx)
}

func (c *ConfigurationService) UpdateFaceRecognitionRegistration(ctx context.Context, registration model.FaceRecognitionRegistration) error {
	return c.repo.UpdateFaceRecognitionRegistration(ctx, registration)
}
