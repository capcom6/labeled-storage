package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"

	"github.com/capcom6/labeled-storage/internal/logger"
	"github.com/capcom6/labeled-storage/internal/repository/memory"
	"github.com/capcom6/labeled-storage/internal/server"
	pb "github.com/capcom6/labeled-storage/pkg/api"
	grpcpkg "github.com/capcom6/labeled-storage/pkg/grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
)

//go:generate protoc --go_out=pkg --go_opt=paths=source_relative --go-grpc_out=pkg --go-grpc_opt=paths=source_relative api/storage.proto

func main() {
	log := logger.Default()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9090))
	if err != nil {
		log.Error("failed to listen", "error", err)
		os.Exit(1)
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(
				grpcpkg.InterceptorLogger(logger.Default()),
			),
		),
	}
	grpcServer := grpc.NewServer(opts...)

	repository := memory.New()

	pb.RegisterStorageServer(grpcServer, server.NewServer(repository))

	wg := sync.WaitGroup{}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		log.Info("shutting down...")
		grpcServer.GracefulStop()
	}()

	log.Info("server started", "address", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Error("failed to serve", "error", err)
		os.Exit(1)
	}
	wg.Wait()

	log.Info("server stopped")
}
