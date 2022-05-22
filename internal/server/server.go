package server

import (
	"fmt"
	"github.com/aveplen-bach/config-service/internal/api"
	pb "github.com/aveplen-bach/config-service/protos/config"
	"google.golang.org/grpc"
	"net"

	"github.com/aveplen-bach/config-service/internal/repo"
)

func Start() {
	db, err := repo.NewRepository()
	if err != nil {
		panic(err)
	}

	server := api.NewGrpcConfigServer(db)

	grpcServer := grpc.NewServer()
	pb.RegisterConfigServer(grpcServer, server)
	listen, _ := net.Listen("tcp", ":8080")

	if err := grpcServer.Serve(listen); err != nil {
		fmt.Println(err)
	}
}
