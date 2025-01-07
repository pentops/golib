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

func Opt[T any](v *T) (o T) {
	if v == nil {
		return // the default value, "", 0, false etc
	}
	return *v
}

func MustUnwrap[T any](v *T) T {
	if v == nil {
		panic("nil value")
	}
	return *v
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
