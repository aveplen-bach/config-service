package server

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aveplen-bach/config-service/internal/config"
	"github.com/aveplen-bach/config-service/internal/repo"
	"github.com/aveplen-bach/config-service/internal/service"
	"github.com/aveplen-bach/config-service/internal/transport"
	pb "github.com/aveplen-bach/config-service/protos/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Start(cfg config.Config) {
	// ================================= repo =================================
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
	if err := grpcServer.Serve(listen); err != nil {
		if cfg.ServerConfig.Debug != nil {
			logrus.Fatal("could not start grpc server: %w", err)
		} else {
			logrus.Warn("could not start grpc server: %w", err)
		}
	}

	// ================================ router ================================
	r := gin.Default()

	// ================================ routes ================================
	r.GET("/ready", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// =============================== shutdown ===============================
	srv := &http.Server{
		Addr:    cfg.ServerConfig.Addr,
		Handler: r,
	}

	go func() {
		logrus.Infof("listening: %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			logrus.Warn(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Warn("shutting down server...")

	ctx, authCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer authCancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Server forced to shutdown:", err)
	}

	logrus.Warn("server exited")
}
