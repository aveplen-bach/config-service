package api

import (
	"context"

	"github.com/aveplen-bach/config-service/internal/model"
	pb "github.com/aveplen-bach/config-service/protos/config"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (c *ConfigServerApi) GetFacerecConfig(ctx context.Context, req *emptypb.Empty) (*pb.FacerecConfig, error) {

	config, err := c.service.GetFacerecConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.FacerecConfig{
		Threshold: config.Threshold,
	}, nil

}

func (c *ConfigServerApi) UpdateFacerecConfig(ctx context.Context, req *pb.FacerecConfig) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, c.service.UpdateFacerecConfig(
		ctx,
		model.FacerecConfig{
			Threshold: req.Threshold,
		})
}
