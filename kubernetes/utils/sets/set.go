package sets

type Empty struct{}
type Set[T comparable] map[T]Empty

func cast[T comparable](s map[T]Empty) Set[T] { return s }
func New[T comparable](items ...T) Set[T] {
	ss := make(Set[T], len(items))
	ss.Insert(items...)
	return ss
}

func (s Set[T]) Insert(items ...T) Set[T] {
	for _, item := range items {
		s[item] = Empty{}
	}
	return s
}

func (s Set[T]) Has(item T) bool {
	_, contained := s[item]
	return contained
}
