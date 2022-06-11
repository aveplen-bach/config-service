package server

import (
	"fmt"
	"log"
	"net"

	"github.com/aveplen-bach/config-service/internal/api"
	"github.com/aveplen-bach/config-service/internal/service"
	pb "github.com/aveplen-bach/config-service/protos/config"
	"google.golang.org/grpc"

	"github.com/aveplen-bach/config-service/internal/repo"
)

func Start() {
	repo, err := repo.NewRepository()
	if err != nil {
		panic(err)
	}

	configurationService := service.NewConfigurationService(repo)

	api := api.NewConfigServerApi(configurationService)

	grpcServer := grpc.NewServer()
	pb.RegisterConfigServer(grpcServer, api)
	listen, _ := net.Listen("tcp", ":30033")

	log.Println("listening on port 30033")
	if err := grpcServer.Serve(listen); err != nil {
		fmt.Println(err)
	}
}
