package memory

import (
	"context"

	"sync"

	"github.com/capcom6/labeled-storage/internal/repository/data"
	"github.com/capcom6/labeled-storage/internal/repository/errors"
	m "github.com/capcom6/labeled-storage/pkg/maps"
	"golang.org/x/exp/maps"
)

type Repository struct {
	keys   map[keyString]data.Item
	labels tree

	mux sync.RWMutex
}

func New() *Repository {
	return &Repository{
		keys:   map[keyString]data.Item{},
		labels: tree{},

		mux: sync.RWMutex{},
	}
}

func (r *Repository) Get(ctx context.Context, key string) (data.Item, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	if item, ok := r.keys[keyString(key)]; ok {
		return item, nil
	}
	return data.Item{}, errors.ErrKeyNotFound
}

func (r *Repository) List(ctx context.Context, labels map[string]string) ([]data.Item, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	if len(labels) == 0 {
		return maps.Values(r.keys), nil
	}

	found := map[keyString]int{}
	for k, v := range labels {
		keys := r.labels.get(keyString(k), v)

		for _, k := range keys {
			found[k]++
		}
	}

	keys := make([]keyString, 0, len(found))
	for k, c := range found {
		if c != len(labels) {
			continue
		}

		keys = append(keys, k)
	}

	return m.SubsetValues(r.keys, keys...), nil
}

func (r *Repository) Replace(ctx context.Context, key string, item data.Item) error {
	panic("not implemented") // TODO: Implement
}

func (r *Repository) DeleteOne(ctx context.Context, key string) error {
	panic("not implemented") // TODO: Implement
}

func (r *Repository) DeleteMany(ctx context.Context, labels map[string]string) error {
	panic("not implemented") // TODO: Implement
}
