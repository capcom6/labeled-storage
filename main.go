package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"

	"github.com/capcom6/labeled-storage/internal/server"
	pb "github.com/capcom6/labeled-storage/pkg/api"
	"google.golang.org/grpc"
)

//go:generate protoc --go_out=pkg --go_opt=paths=source_relative --go-grpc_out=pkg --go-grpc_opt=paths=source_relative api/storage.proto

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterStorageServer(grpcServer, server.NewServer())

	wg := sync.WaitGroup{}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		log.Println("shutting down...")
		grpcServer.GracefulStop()
	}()

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	wg.Wait()

	log.Println("server stopped")
}
