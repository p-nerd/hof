package hof

import "errors"

func ForEach[T any](slice []T, callback func(int, T, []T)) error {
	if slice == nil {
		return errors.New("slice is nil")
	}
	if callback == nil {
		return errors.New("callback is nil")
	}
	for i, item := range slice {
		callback(i, item, slice)
	}
	return nil
}

func Map[T, U any](slice []T, callback func(T) U) ([]U, error) {
	if slice == nil {
		return nil, errors.New("slice is nil")
	}
	if callback == nil {
		return nil, errors.New("callback is nil")
	}

	mapped := make([]U, len(slice))
	for i, v := range slice {
		mapped[i] = callback(v)
	}
	return mapped, nil
}
