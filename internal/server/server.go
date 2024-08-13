package server

import (
	"context"

	"github.com/capcom6/go-helpers/slices"
	"github.com/capcom6/labeled-storage/internal/repository"
	"github.com/capcom6/labeled-storage/internal/repository/data"
	pb "github.com/capcom6/labeled-storage/pkg/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedStorageServer

	repo repository.Interface
}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {

	item, err := s.repo.Get(ctx, req.GetKey())
	if err != nil {
		if err == repository.ErrKeyNotFound {
			return nil, status.Errorf(codes.NotFound, "key %s not found", req.GetKey())
		}

		return nil, err
	}

	return &pb.GetResponse{Item: &pb.Item{Key: item.Key(), Value: item.Value(), Labels: item.Labels()}}, nil
}
func (s *Server) Find(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error) {
	items, err := s.repo.List(ctx, req.GetLabels())
	if err != nil {
		return nil, err
	}

	return &pb.FindResponse{
		Items: slices.Map(items, func(item data.ItemWithKey) *pb.Item {
			return &pb.Item{Key: item.Key(), Value: item.Value(), Labels: item.Labels()}
		}),
	}, nil
}
func (s *Server) Replace(ctx context.Context, req *pb.ReplaceRequest) (*pb.ReplaceResponse, error) {
	item, err := s.repo.Replace(ctx, req.Item.Key, *data.NewItem(req.Item.Value, req.Item.Labels))
	if err != nil {
		return nil, err
	}

	return &pb.ReplaceResponse{Item: &pb.Item{Key: item.Key(), Value: item.Value(), Labels: item.Labels()}}, nil
}
func (s *Server) DeleteOne(ctx context.Context, req *pb.DeleteOneRequest) (*pb.DeleteOneResponse, error) {
	if req.Key == "" {
		return nil, status.Errorf(codes.InvalidArgument, "key is required")
	}

	if err := s.repo.DeleteOne(ctx, req.Key); err != nil {
		return nil, err
	}

	return &pb.DeleteOneResponse{}, nil
}

func (s *Server) DeleteMany(ctx context.Context, req *pb.DeleteManyRequest) (*pb.DeleteManyResponse, error) {
	cnt, err := s.repo.DeleteMany(ctx, req.Labels)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteManyResponse{
		Deleted: uint64(cnt),
	}, nil
}

func NewServer(repo repository.Interface) *Server {
	if repo == nil {
		panic("repo cannot be nil")
	}

	return &Server{
		repo: repo,
	}
}
