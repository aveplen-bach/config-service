package transport

import (
	"context"

	"github.com/aveplen-bach/config-service/internal/model"
	"github.com/aveplen-bach/config-service/internal/service"
	pb "github.com/aveplen-bach/config-service/protos/config"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	pb.UnimplementedConfigServer
	service *service.ConfigurationService
}

func NewService(service *service.ConfigurationService) *Service {
	return &Service{
		service: service,
	}
}

func (c *Service) GetFacerecConfig(ctx context.Context, req *emptypb.Empty) (*pb.FacerecConfig, error) {

	config, err := c.service.GetFacerecConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.FacerecConfig{
		Threshold: config.Threshold,
	}, nil

}

func (c *Service) UpdateFacerecConfig(ctx context.Context, req *pb.FacerecConfig) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, c.service.UpdateFacerecConfig(
		ctx,
		model.FacerecConfig{
			Threshold: req.Threshold,
		})
}
