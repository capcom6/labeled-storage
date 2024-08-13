package repository

import (
	"context"

	"github.com/capcom6/labeled-storage/internal/repository/data"
	"github.com/capcom6/labeled-storage/internal/repository/errors"
)

const ErrKeyNotFound = errors.ErrKeyNotFound

type Interface interface {
	Get(ctx context.Context, key string) (data.ItemWithKey, error)
	List(ctx context.Context, labels map[string]string) ([]data.ItemWithKey, error)
	Replace(ctx context.Context, key string, item data.Item) (data.ItemWithKey, error)
	DeleteOne(ctx context.Context, key string) error
	DeleteMany(ctx context.Context, labels map[string]string) (int, error)
}
