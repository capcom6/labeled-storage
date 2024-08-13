package repository

import (
	"context"

	"github.com/capcom6/labeled-storage/internal/repository/data"
)

type Interface interface {
	Get(ctx context.Context, key string) (data.Item, error)
	List(ctx context.Context, labels map[string]string) ([]data.Item, error)
	Replace(ctx context.Context, key string, item data.Item) error
	DeleteOne(ctx context.Context, key string) error
	DeleteMany(ctx context.Context, labels map[string]string) error
}
