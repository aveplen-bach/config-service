package server

import (
	"net"

	"github.com/aveplen-bach/config-service/internal/config"
	"github.com/aveplen-bach/config-service/internal/repo"
	"github.com/aveplen-bach/config-service/internal/service"
	"github.com/aveplen-bach/config-service/internal/transport"
	pb "github.com/aveplen-bach/config-service/protos/config"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Start(cfg config.Config) {
	// ================================= repo =================================
	logrus.Info("creating repository")
	repo, err := repo.NewRepository()
	if err != nil {
		logrus.Fatal("could not connect to db: %w", err)
	}

	// ================================ server ================================
	configServer := service.NewConfigurationService(repo)

	facerec := transport.NewService(configServer)
	grpcServer := grpc.NewServer()
	pb.RegisterConfigServer(grpcServer, facerec)
	listen, _ := net.Listen("tcp", cfg.ServerConfig.Addr)

	logrus.Infof("listening on %s", cfg.ServerConfig.Addr)
	if err := grpcServer.Serve(listen); err != nil {
		logrus.Warn("server failed: %w", err)
	}
}
