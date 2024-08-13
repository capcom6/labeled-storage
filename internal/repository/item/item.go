package item

import (
	"fmt"
	"maps"
)

type Item struct {
	value  string
	labels map[string]string
}

func New(value string, labels map[string]string) *Item {
	return &Item{
		value:  value,
		labels: labels,
	}
}

func (i *Item) Value() string {
	return i.value
}

func (i *Item) Labels() map[string]string {
	return maps.Clone(i.labels)
}

func (i *Item) String() string {
	return fmt.Sprintf("key: %s, value: %s, labels: %v", i.value, i.value, i.labels)
}
