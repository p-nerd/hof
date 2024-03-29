package hof_test

import (
	"reflect"
	"testing"

	"github.com/p-nerd/hof"
)

// Unhappy case

func TestFilterNilSlice(t *testing.T) {
	filteredSlice := hof.Filter(nil, func(_ int, item int, _ []int) bool {
		return true
	})
	if filteredSlice == nil {
		t.Errorf("Expected filtered slice to be emty slice, got: %v", filteredSlice)
	}
}

func TestFilterNilCallback(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	filteredSlice := hof.Filter(intSlice, nil)
	if filteredSlice == nil {
		t.Errorf(
			"Expected filtered slice to be emty slice, got: %v",
			filteredSlice,
		)
	}
}

// Happy case

func TestFilterIntSlice(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	filteredSlice := hof.Filter(intSlice, func(_ int, item int, _ []int) bool {
		return item%2 == 0
	})
	expected := []int{2, 4}
	if !reflect.DeepEqual(filteredSlice, expected) {
		t.Errorf(
			"Filtered slice does not match the expected result. Expected: %v, Got: %v",
			expected,
			filteredSlice,
		)
	}
}
