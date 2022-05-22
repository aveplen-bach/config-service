package api

import (
	"github.com/aveplen-bach/config-service/internal/repo"
	pb "github.com/aveplen-bach/config-service/protos/config"
)

type GrpcConfigServer struct {
	pb.UnimplementedConfigServer
	repo *repo.Repository
}

func NewGrpcConfigServer(repo *repo.Repository) *GrpcConfigServer {
	return &GrpcConfigServer{
		repo: repo,
	}
}
