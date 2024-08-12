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

func (s *Server) Get(context.Context, *pb.GetRequest) (*pb.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (s *Server) List(context.Context, *pb.ListRequest) (*pb.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (s *Server) Put(context.Context, *pb.PutRequest) (*pb.PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (s *Server) Delete(context.Context, *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func NewServer() *Server {
	return &Server{}
}
