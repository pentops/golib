package gl

//go:fix inline
func Ptr[T any](v T) *T {
	return new(v)
}

// Deprecated
func Coalesce[T any](fallback T, try ...*T) T {
	for _, v := range try {
		if v != nil {
			return *v
		}
	}
	return fallback
}

// Deprecated
func Opt[T any](v *T) (o T) {
	if v == nil {
		return // the default value, "", 0, false etc
	}
	return *v
}

// Deprecated
func MustUnwrap[T any](v *T) T {
	if v == nil {
		panic("nil value")
	}
	return *v
}

// Deprecated
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
