package maps

import (
	"golang.org/x/exp/maps"
)

func Subset[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	result := make(map[K]V, len(keys))

	for _, k := range keys {
		if v, ok := m[k]; ok {
			result[k] = v
		}
	}

	return result
}

func SubsetValues[K comparable, V any](m map[K]V, keys ...K) []V {
	return maps.Values(Subset(m, keys...))
}
