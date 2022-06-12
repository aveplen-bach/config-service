package service

import "github.com/aveplen-bach/config-service/internal/repo"

type ConfigurationService struct {
	repo *repo.Repository
}

func NewConfigurationService(repo *repo.Repository) *ConfigurationService {
	return &ConfigurationService{repo: repo}
}
