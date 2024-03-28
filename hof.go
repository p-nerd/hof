// This package provides higher-order functions for working with slices.
package hof

import "errors"

// ForEach applies a callback function to each element in the input slice.
//
// The callback function, callback, should accept three parameters:
//   - The index of the element in the slice.
//   - The element itself.
//   - The entire input slice.
//
// The callback function does not return any value.
//
// Example:
//
//	// Print each element in a slice of strings
//	err := ForEach([]string{"apple", "banana", "cherry"}, func(i int, item string, slice []string) {
//	    fmt.Println(i, item)
//	})
//
// Parameters:
//   - slice: The input slice to iterate over.
//   - callback: The callback function to be applied to each element.
//
// Returns:
//   - error: An error if the input slice or callback function is nil.
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

// Map applies a mapping function to each element in the input slice
// and returns a new slice containing the results of applying the
// mapping function to each element.
//
// The mapping function, callback, should accept three parameters:
//   - The index of the element in the slice.
//   - The element itself.
//   - The entire input slice.
//
// The mapping function should return the transformed element of type U.
//
// Example:
//
//	// Double each element in a slice of integers
//	mappedSlice, err := Map([]int{1, 2, 3, 4, 5}, func(i int, x int, slice []int) int {
//	    return x * 2
//	})
//
// Parameters:
//   - slice: The input slice to be mapped.
//   - callback: The mapping function to be applied to each element.
//
// Returns:
//   - []U: A new slice containing the results of applying the mapping function to each element.
//   - error: An error if the input slice or callback function is nil.
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

// Filter applies a filtering function to each element in the input slice
// and returns a new slice containing only the elements for which the
// filtering function returns true.
//
// The filtering function, callback, should accept an element of the slice
// and return a boolean value indicating whether the element should be
// included in the filtered result.
//
// Example:
//
//	// Filter out even numbers from a slice of integers
//	filteredSlice, err := Filter([]int{1, 2, 3, 4, 5}, func(x int) bool {
//	    return x%2 != 0
//	})
//
// Parameters:
//   - slice: The input slice to be filtered.
//   - callback: The filtering function to be applied to each element.
//
// Returns:
//   - []T: A new slice containing only the elements that satisfy the filtering condition.
//   - error: An error if the input slice or callback function is nil.
func Filter[T any](slice []T, callback func(int, T, []T) bool) ([]T, error) {
	if slice == nil {
		return nil, errors.New("slice is nil")
	}
	if callback == nil {
		return nil, errors.New("callback is nil")
	}
	filteredSlice := []T{}
	for index, item := range slice {
		if callback(index, item, slice) {
			filteredSlice = append(filteredSlice, item)
		}
	}
	return filteredSlice, nil
}
