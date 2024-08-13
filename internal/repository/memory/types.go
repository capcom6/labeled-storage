package memory

type keyString string

type leaf map[string][]keyString

func (l leaf) get(key string) []keyString {
	return l[key]
}

type tree map[keyString]leaf

func (t tree) get(key keyString, value string) []keyString {
	return t[key].get(value)
}
