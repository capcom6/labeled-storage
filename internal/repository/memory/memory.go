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
	keys   map[keyString]data.ItemWithKey
	labels tree

	mux sync.RWMutex
}

func New() *Repository {
	return &Repository{
		keys:   map[keyString]data.ItemWithKey{},
		labels: tree{},

		mux: sync.RWMutex{},
	}
}

func (r *Repository) Get(ctx context.Context, key string) (data.ItemWithKey, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	if item, ok := r.keys[keyString(key)]; ok {
		return *data.NewItemWithKey(key, item.Value(), item.Labels()), nil
	}
	return data.ItemWithKey{}, errors.ErrKeyNotFound
}

func (r *Repository) List(ctx context.Context, labels map[string]string) ([]data.ItemWithKey, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	if len(labels) == 0 {
		return maps.Values(r.keys), nil
	}

	keys := r.findKeys(labels)

	return m.SubsetValues(r.keys, keys...), nil
}

func (r *Repository) Replace(ctx context.Context, key string, item data.Item) (data.ItemWithKey, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.delete(keyString(key))

	r.keys[keyString(key)] = *data.NewItemWithKey(key, item.Value(), item.Labels())
	for l, v := range item.Labels() {
		r.labels.add(l, v, keyString(key))
	}
	return r.keys[keyString(key)], nil
}

func (r *Repository) DeleteOne(ctx context.Context, key string) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	if cnt := r.delete(keyString(key)); cnt == 0 {
		return errors.ErrKeyNotFound
	}

	return nil
}

func (r *Repository) DeleteMany(ctx context.Context, labels map[string]string) (int, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	keys := r.findKeys(labels)

	cnt := r.delete(keys...)

	return cnt, nil
}

func (r *Repository) delete(keys ...keyString) int {
	cnt := 0
	for _, key := range keys {
		item, ok := r.keys[keyString(key)]
		if !ok {
			continue
		}

		for l, v := range item.Labels() {
			r.labels.remove(l, v, keyString(key))
		}

		delete(r.keys, keyString(key))
		cnt++
	}
	return cnt
}

func (r *Repository) findKeys(labels map[string]string) []keyString {
	found := map[keyString]int{}
	for l, v := range labels {
		keys := r.labels.get(l, v)

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

	return keys
}
