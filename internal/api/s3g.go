package api

import (
	"context"
	"github.com/aveplen-bach/config-service/internal/model"
	pb "github.com/aveplen-bach/config-service/protos/config"
)

func (c *ConfigServerApi) GetS3GatewayConfig(ctx context.Context, req *pb.Ack) (*pb.S3GatewayConfig, error) {
	config, err := c.service.GetS3GatewayConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.S3GatewayConfig{
		Host: config.Host,
		Port: config.Port,
	}, nil
}

func (c *ConfigServerApi) UpdateS3GatewayConfig(ctx context.Context, req *pb.S3GatewayConfig) (*pb.Ack, error) {
	return &pb.Ack{}, c.service.UpdateS3GatewayConfig(
		ctx,
		model.S3GatewayConfig{
			Host: req.Host,
			Port: req.Port,
		})
}

func (c *ConfigServerApi) RegisterS3Gateway(ctx context.Context, req *pb.S3GatewayRegistration) (*pb.Ack, error) {
	return &pb.Ack{}, c.service.UpdateS3GatewayRegistration(
		ctx,
		model.S3GatewayRegistration{
			Host: req.Host,
			Port: req.Port,
		})
}
