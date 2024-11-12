package gl

func Ptr[T any](v T) *T {
	return &v
}

func Coalesce[T any](fallback T, try ...*T) T {
	for _, v := range try {
		if v != nil {
			return *v
		}
	}
	return fallback
}
