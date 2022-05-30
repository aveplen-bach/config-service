package service

import (
	"context"
	"github.com/aveplen-bach/config-service/internal/model"
	"log"
)

func (c *ConfigurationService) GetS3GatewayConfig(ctx context.Context) (model.S3GatewayConfig, error) {
	log.Println("[ConfigurationService] GetS3GatewayConfig")
	return c.repo.GetS3GatewayConfig(ctx)
}

func (c *ConfigurationService) UpdateS3GatewayConfig(ctx context.Context, config model.S3GatewayConfig) error {
	log.Println("[ConfigurationService] UpdateS3GatewayConfig")
	return c.repo.UpdateS3GatewayConfigConfig(ctx, config)
}

func (c *ConfigurationService) GetS3GatewayRegistration(ctx context.Context) (model.S3GatewayRegistration, error) {
	log.Println("[ConfigurationService] GetS3GatewayRegistration")
	return c.repo.GetS3GatewayRegistration(ctx)
}

func (c *ConfigurationService) UpdateS3GatewayRegistration(ctx context.Context, registration model.S3GatewayRegistration) error {
	log.Println("[ConfigurationService] UpdateS3GatewayRegistration")
	return c.repo.UpdateS3GatewayRegistration(ctx, registration)
}
