package data

import (
	"fmt"
	"maps"
)

type Item struct {
	value  string
	labels map[string]string
}

func NewItem(value string, labels map[string]string) *Item {
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
	return fmt.Sprintf("value: %s, labels: %v", i.value, i.labels)
}

type ItemWithKey struct {
	key string
	Item
}

func NewItemWithKey(key string, value string, labels map[string]string) *ItemWithKey {
	return &ItemWithKey{
		key:  key,
		Item: *NewItem(value, labels),
	}
}

func (i *ItemWithKey) Key() string {
	return i.key
}

func (i *ItemWithKey) String() string {
	return fmt.Sprintf("key: %s, value: %s, labels: %v", i.key, i.value, i.labels)
}
