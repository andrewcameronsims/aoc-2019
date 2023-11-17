package set

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(map[T]struct{})
}

func (s Set[T]) Add(t T) {
	s[t] = struct{}{}
}

func (s Set[T]) AddAll(ts ...T) {
	for _, t := range ts {
		s.Add(t)
	}
}

func (s Set[T]) Remove(t T) {
	delete(s, t)
}

func (s Set[T]) Has(t T) bool {
	_, ok := s[t]
	return ok
}

func (s Set[T]) Intersect(other Set[T]) Set[T] {
	result := NewSet[T]()

	for t := range s {
		if other.Has(t) {
			result.Add(t)
		}
	}

	return result
}
