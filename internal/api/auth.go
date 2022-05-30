package api

import (
	"context"
	"github.com/aveplen-bach/config-service/internal/model"
	pb "github.com/aveplen-bach/config-service/protos/config"
)

func (c *ConfigServerApi) GetAuthenticationConfig(ctx context.Context, req *pb.Ack) (*pb.AuthenticationConfig, error) {
	config, err := c.service.GetAuthenticationConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.AuthenticationConfig{
		DatabaseConfig: &pb.DatabaseConfig{
			Host:     config.DatabaseConfig.Host,
			Port:     config.DatabaseConfig.Port,
			Username: config.DatabaseConfig.Username,
			Password: config.DatabaseConfig.Password,
			Database: config.DatabaseConfig.Database,
		},
		FacerecRegistration: &pb.FaceRecognitionRegistration{
			Host: config.FacerecRegistration.Host,
			Port: config.FacerecRegistration.Port,
		},
		S3GatewayRegistration: &pb.S3GatewayRegistration{
			Host: config.S3GatewayRegistration.Host,
			Port: config.S3GatewayRegistration.Port,
		},
	}, nil
}

func (c *ConfigServerApi) UpdateAuthenticationDatabaseConfig(ctx context.Context, req *pb.DatabaseConfig) (*pb.Ack, error) {
	return &pb.Ack{}, c.service.UpdateAuthenticationDatabaseConfig(ctx, model.AuthenticationDatabaseConfig{
		Host:     req.Host,
		Port:     req.Port,
		Username: req.Username,
		Password: req.Password,
		Database: req.Database,
	})
}

func (c *ConfigServerApi) RegisterAuthentication(ctx context.Context, req *pb.AuthenticationRegistration) (*pb.Ack, error) {
	return &pb.Ack{}, c.service.UpdateAuthenticationRegistration(
		ctx,
		model.AuthenticationRegistration{
			Host: req.Host,
			Port: req.Port,
		})
}
