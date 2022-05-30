package api

import (
	"context"
	pb "github.com/aveplen-bach/config-service/protos/config"
)

func (c *ConfigServerApi) GetAuthenticationClientConfig(ctx context.Context, req *pb.Ack) (*pb.AuthenticationClientConfig, error) {
	config, err := c.service.GetAuthenticationClientConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.AuthenticationClientConfig{
		Registration: &pb.AuthenticationRegistration{
			Host: config.Registration.Host,
			Port: config.Registration.Port,
		}}, nil
}
