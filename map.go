package hof

func Map[T, U any](arr []T, f func(T) U) []U {
	mapped := make([]U, len(arr))
	for i, v := range arr {
		mapped[i] = f(v)
	}
	return mapped
}
