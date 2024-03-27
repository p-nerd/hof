// This package provides higher-order functions for working with slices.
package hof

import "errors"

// ForEach applies a function to each element of a slice.
//
// The function callback is called with the index of the element,
// the element itself, and the entire slice.
//
// If slice is nil, it returns an error.
//
// If callback is nil, it returns an error.
func ForEach[T any](slice []T, callback func(int, T, []T)) error {
	if slice == nil {
		return errors.New("slice is nil")
	}
	if callback == nil {
		return errors.New("callback is nil")
	}
	for index, item := range slice {
		callback(index, item, slice)
	}
	return nil
}

// Map applies a function to each element of a slice and returns a new slice
// containing the results of applying the function to each element.
//
// The function callback is called with the index of the element,
// the element itself, and the entire slice.
//
// If slice is nil, it returns nil and an error.
//
// If callback is nil, it returns nil and an error.
func Map[T, U any](slice []T, callback func(int, T, []T) U) ([]U, error) {
	if slice == nil {
		return nil, errors.New("slice is nil")
	}
	if callback == nil {
		return nil, errors.New("callback is nil")
	}
	mapped := make([]U, len(slice))
	for index, item := range slice {
		mapped[index] = callback(index, item, slice)
	}
	return mapped, nil
}
