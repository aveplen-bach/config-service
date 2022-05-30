package service

import (
	"context"
	"github.com/aveplen-bach/config-service/internal/model"
)

func (c *ConfigurationService) GetAuthenticationClientConfig(ctx context.Context) (model.AuthenticationClientConfig, error) {
	authRegistration, err := c.GetAuthenticationRegistration(ctx)
	if err != nil {
		return model.AuthenticationClientConfig{}, err
	}

	return model.AuthenticationClientConfig{
		Registration: authRegistration,
	}, nil
}
