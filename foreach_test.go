package hof_test

import (
	"testing"

	"github.com/p-nerd/hof"
)

// Unhappy case

func TestForEachNilSlice(t *testing.T) {
	err := hof.ForEach(nil, func(i int, item int, slice []int) {
		t.Error("Function should not be called for nil slice")
	})
	if err == nil {
		t.Error("Expected error for nil slice, got nil")
	} else if err.Error() != "slice is nil" {
		t.Errorf("Expected error message 'slice is nil', got '%s'", err.Error())
	}
}

func TestForEachNilCallback(t *testing.T) {
	err := hof.ForEach([]int{1, 2, 3}, nil)
	if err == nil {
		t.Error("Expected error for nil function, got nil")
	} else if err.Error() != "callback is nil" {
		t.Errorf("Expected error message 'callback is nil', got '%s'", err.Error())
	}
}

// Happy case

func TestForEachEmptySlice(t *testing.T) {
	err := hof.ForEach([]int{}, func(i int, item int, slice []int) {
		t.Error("Function should not be called for empty slice")
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestForEachValidSlice(t *testing.T) {
	validSlice := []int{1}
	var called bool
	err := hof.ForEach(validSlice, func(i int, item int, slice []int) {
		called = true
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !called {
		t.Error("Function was not called for non-empty slice")
	}
}

// Boolean
func TestForEachBoolSlice(t *testing.T) {
	boolSlice := []bool{true, false, true}
	err := hof.ForEach(boolSlice, func(i int, item bool, slice []bool) {
		if slice[i] != item {
			t.Errorf("Item is not matching: index=%d, %v", i, item)
		}
		if i < 0 || i >= len(slice) {
			t.Errorf("Index out of range: index=%d, len=%d", i, len(slice))
		}
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

// String
func TestForEachStringSlice(t *testing.T) {
	stringSlice := []string{"apple", "banana", "cherry"}
	err := hof.ForEach(stringSlice, func(i int, item string, slice []string) {
		if slice[i] != item {
			t.Errorf("Item is not matching: index=%d, %v", i, item)
		}
		if i < 0 || i >= len(slice) {
			t.Errorf("Index out of range: index=%d, len=%d", i, len(slice))
		}
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

// Numeric
func TestForEachIntSlice(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	err := hof.ForEach(intSlice, func(i int, item int, slice []int) {
		if slice[i] != item {
			t.Errorf("Item is not matching: index=%d, %v", i, item)
		}
		if i < 0 || i >= len(slice) {
			t.Errorf("Index out of range: index=%d, len=%d", i, len(slice))
		}
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestForEachUintSlice(t *testing.T) {
	uintSlice := []uint{1, 2, 3, 4, 5}
	err := hof.ForEach(uintSlice, func(i int, item uint, slice []uint) {
		if slice[i] != item {
			t.Errorf("Item is not matching: index=%d, %v", i, item)
		}
		if i < 0 || i >= len(slice) {
			t.Errorf("Index out of range: index=%d, len=%d", i, len(slice))
		}
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestForEachUintptrSlice(t *testing.T) {
	uintptrSlice := []uintptr{1, 2, 3, 4, 5}
	err := hof.ForEach(uintptrSlice, func(i int, item uintptr, slice []uintptr) {
		if slice[i] != item {
			t.Errorf("Item is not matching: index=%d, %v", i, item)
		}
		if i < 0 || i >= len(slice) {
			t.Errorf("Index out of range: index=%d, len=%d", i, len(slice))
		}
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

}

func TestForEachFloat64Slice(t *testing.T) {
	float64Slice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	err := hof.ForEach(float64Slice, func(i int, item float64, slice []float64) {
		if slice[i] != item {
			t.Errorf("Item is not matching: index=%d, %v", i, item)
		}
		if i < 0 || i >= len(slice) {
			t.Errorf("Index out of range: index=%d, len=%d", i, len(slice))
		}
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
