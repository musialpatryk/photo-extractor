package organizer

type Entry interface {
	GroupKey() string
}

type Bucket[T Entry] struct {
	Key   string
	Items []T
}

func Organize[T Entry](entries []T) []Bucket[T] {
	groups := make(map[string][]T)

	for _, entry := range entries {
		key := entry.GroupKey()
		groups[key] = append(groups[key], entry)
	}

	buckets := make([]Bucket[T], 0, len(groups))
	for key, items := range groups {
		buckets = append(buckets, Bucket[T]{
			Key:   key,
			Items: items,
		})
	}

	return buckets
}
