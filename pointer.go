package pointer

func Of[T any](v T) *T {
	return &v
}

func ValueOrDefault[T any](p *T, _default T) T {
	if p == nil {
		return _default
	}

	return *p
}

func ValueOrEmpty[T any](p *T) T {
	return ValueOrDefault(p, *new(T))
}
