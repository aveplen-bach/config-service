package service

import (
	"context"
	"github.com/aveplen-bach/config-service/internal/model"
)

func (c *ConfigurationService) GetAuthenticationConfig(ctx context.Context) (model.AuthenticationConfig, error) {
	authDBConfig, err := c.GetAuthenticationDatabaseConfig(ctx)
	if err != nil {
		return model.AuthenticationConfig{}, err
	}

	facerecRegistration, err := c.GetFaceRecognitionRegistration(ctx)
	if err != nil {
		return model.AuthenticationConfig{}, err
	}

	s3gRegistration, err := c.GetS3GatewayRegistration(ctx)
	if err != nil {
		return model.AuthenticationConfig{}, err
	}

	return model.AuthenticationConfig{
		DatabaseConfig:        authDBConfig,
		FacerecRegistration:   facerecRegistration,
		S3GatewayRegistration: s3gRegistration,
	}, nil
}

func (c *ConfigurationService) GetAuthenticationDatabaseConfig(ctx context.Context) (model.AuthenticationDatabaseConfig, error) {
	return c.repo.GetAuthenticationDatabaseConfig(ctx)
}

func (c *ConfigurationService) UpdateAuthenticationDatabaseConfig(ctx context.Context, config model.AuthenticationDatabaseConfig) error {
	return c.repo.UpdateAuthenticationDatabaseConfig(ctx, config)
}

func (c *ConfigurationService) GetAuthenticationRegistration(ctx context.Context) (model.AuthenticationRegistration, error) {
	return c.repo.GetAuthenticationRegistration(ctx)
}

func (c *ConfigurationService) UpdateAuthenticationRegistration(ctx context.Context, registration model.AuthenticationRegistration) error {
	return c.repo.UpdateAuthenticationRegistration(ctx, registration)
}
