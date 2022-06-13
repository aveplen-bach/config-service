package service

import (
	"context"
	"fmt"

	"github.com/aveplen-bach/config-service/internal/model"
	"github.com/sirupsen/logrus"
)

func (c *ConfigurationService) GetFacerecConfig(ctx context.Context) (model.FacerecConfig, error) {
	logrus.Info("getting facerec config")
	config, err := c.repo.GetFaceRecognitionConfig(ctx)
	if err != nil {
		logrus.Warnf("could not get facerec config: %w", err)
		return model.FacerecConfig{}, fmt.Errorf("could not get facerec config: %w", err)
	}
	return config, nil
}

func (c *ConfigurationService) UpdateFacerecConfig(ctx context.Context, config model.FacerecConfig) error {
	if err := c.repo.UpdateFaceRecognitionConfig(ctx, config); err != nil {
		logrus.Warnf("could not update facerec config: %w", err)
		return fmt.Errorf("could not update facerec config: %w", err)
	}
	return nil
}
