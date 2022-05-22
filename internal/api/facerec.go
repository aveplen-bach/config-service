package api

import (
	"context"
	pb "github.com/aveplen-bach/config-service/protos/config"
)

func (g *GrpcConfigServer) GetFaceRecognitionConfig(
	ctx context.Context,
	req *pb.Syn,
) (*pb.FaceRecognitionConfig, error) {

	config, err := g.repo.GetFaceRecognitionConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.FaceRecognitionConfig{
		Threshold: config.Threshold,
	}, nil
}

func (g *GrpcConfigServer) UpdateFaceRecognitionConfig(
	ctx context.Context,
	req *pb.FaceRecognitionConfig,
) (*pb.Ack, error) {

	return nil, nil
}
