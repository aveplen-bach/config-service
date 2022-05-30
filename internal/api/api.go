package api

import (
	"github.com/aveplen-bach/config-service/internal/service"
	pb "github.com/aveplen-bach/config-service/protos/config"
)

type ConfigServerApi struct {
	pb.UnimplementedConfigServer
	service *service.ConfigurationService
}

func NewConfigServerApi(service *service.ConfigurationService) *ConfigServerApi {
	return &ConfigServerApi{
		service: service,
	}
}
