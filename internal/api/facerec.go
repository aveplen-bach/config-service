package api

import (
	"context"
	"github.com/aveplen-bach/config-service/internal/model"
	pb "github.com/aveplen-bach/config-service/protos/config"
)

func (c *ConfigServerApi) GetFaceRecognitionConfig(ctx context.Context, req *pb.Ack) (*pb.FaceRecognitionConfig, error) {

	config, err := c.service.GetFaceRecognitionConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.FaceRecognitionConfig{
		Threshold: config.Threshold,
	}, nil

}

func (c *ConfigServerApi) UpdateFaceRecognitionConfig(ctx context.Context, req *pb.FaceRecognitionConfig) (*pb.Ack, error) {
	return &pb.Ack{}, c.service.UpdateFaceRecognitionConfig(
		ctx,
		model.FaceRecognitionConfig{
			Threshold: req.Threshold,
		})
}

func (c *ConfigServerApi) RegisterFaceRecognition(ctx context.Context, req *pb.FaceRecognitionRegistration) (*pb.Ack, error) {
	return &pb.Ack{}, c.service.UpdateFaceRecognitionRegistration(
		ctx,
		model.FaceRecognitionRegistration{
			Host: req.Host,
			Port: req.Port,
		})
}
