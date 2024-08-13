package memory

import "golang.org/x/exp/maps"

type keyString string

type keysCollection map[keyString]struct{}

func (k keysCollection) add(key keyString) {
	k[key] = struct{}{}
}

func (k keysCollection) get() []keyString {
	return maps.Keys(k)
}

func (k keysCollection) remove(key keyString) {
	delete(k, key)
}

type leaf map[string]keysCollection

func (l leaf) add(value string, key keyString) {
	if l[value] == nil {
		l[value] = keysCollection{}
	}

	l[value].add(key)
}

func (l leaf) get(key string) []keyString {
	return l[key].get()
}

func (l leaf) remove(value string, key keyString) {
	if l[value] == nil {
		return
	}

	l[value].remove(key)
}

type tree map[string]leaf

func (t tree) add(label, value string, key keyString) {
	if t[label] == nil {
		t[label] = leaf{}
	}

	t[label].add(value, key)
}

func (t tree) get(label, value string) []keyString {
	if t[label] == nil {
		return nil
	}

	return t[label].get(value)
}

func (t tree) remove(label, value string, key keyString) {
	if t[label] == nil {
		return
	}

	t[label].remove(value, key)
}
