package hof_test

import (
	"reflect"
	"testing"

	"github.com/p-nerd/hof"
)

// Unhappy case

func TestFilterNilSlice(t *testing.T) {
	filteredSlice, err := hof.Filter(nil, func(_ int, item int, _ []int) bool {
		return true
	})
	if err == nil {
		t.Error("Expected error for nil slice, got nil")
	} else if err.Error() != "slice is nil" {
		t.Errorf("Expected error message 'slice is nil', got '%s'", err.Error())
	}
	if filteredSlice != nil {
		t.Errorf("Expected filtered slice to be nil for nil slice input, got: %v", filteredSlice)
	}
}

func TestFilterNilCallback(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	filteredSlice, err := hof.Filter(intSlice, nil)
	if err == nil {
		t.Error("Expected error for nil callback function, got nil")
	} else if err.Error() != "callback is nil" {
		t.Errorf("Expected error message 'callback is nil', got '%s'", err.Error())
	}
	if filteredSlice != nil {
		t.Errorf(
			"Expected filtered slice to be nil for nil callback function, got: %v",
			filteredSlice,
		)
	}
}

// Happy case

func TestFilterIntSlice(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	filteredSlice, err := hof.Filter(intSlice, func(_ int, item int, _ []int) bool {
		return item%2 == 0
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := []int{2, 4}
	if !reflect.DeepEqual(filteredSlice, expected) {
		t.Errorf(
			"Filtered slice does not match the expected result. Expected: %v, Got: %v",
			expected,
			filteredSlice,
		)
	}
}
