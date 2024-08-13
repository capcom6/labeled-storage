package server

import (
	"context"

	pb "github.com/capcom6/labeled-storage/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedStorageServer
}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	// return &pb.GetResponse{
	// 	Item: &pb.Item{
	// 		Key:    req.GetKey(),
	// 		Value:  "Some Value",
	// 		Labels: map[string]string{},
	// 	},
	// }, nil

	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (s *Server) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (s *Server) Replace(ctx context.Context, req *pb.ReplaceRequest) (*pb.ReplaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Replace not implemented")
}
func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func NewServer() *Server {
	return &Server{}
}
