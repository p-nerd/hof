package hof_test

import (
	"testing"

	"github.com/p-nerd/hof"
)

// Unhappy case

func TestForEachNilSlice(t *testing.T) {
	hof.ForEach(nil, func(i int, item int, slice []int) {
		t.Error("Function should not be called for nil slice")
	})
}

func TestForEachNilCallback(t *testing.T) {
	hof.ForEach([]int{1, 2, 3}, nil)
}

// Happy case

func TestForEachEmptySlice(t *testing.T) {
	hof.ForEach([]int{}, func(i int, item int, slice []int) {
		t.Error("Function should not be called for empty slice")
	})
}

func TestForEachValidSlice(t *testing.T) {
	validSlice := []int{1}
	var called bool
	hof.ForEach(validSlice, func(i int, item int, slice []int) {
		called = true
	})
	if !called {
		t.Error("Function was not called for non-empty slice")
	}
}

// Boolean
func TestForEachBoolSlice(t *testing.T) {
	boolSlice := []bool{true, false, true}
	hof.ForEach(boolSlice, func(i int, item bool, slice []bool) {
		if slice[i] != item {
			t.Errorf("Item is not matching: index=%d, %v", i, item)
		}
		if i < 0 || i >= len(slice) {
			t.Errorf("Index out of range: index=%d, len=%d", i, len(slice))
		}
	})
}

// String
func TestForEachStringSlice(t *testing.T) {
	stringSlice := []string{"apple", "banana", "cherry"}
	hof.ForEach(stringSlice, func(i int, item string, slice []string) {
		if slice[i] != item {
			t.Errorf("Item is not matching: index=%d, %v", i, item)
		}
		if i < 0 || i >= len(slice) {
			t.Errorf("Index out of range: index=%d, len=%d", i, len(slice))
		}
	})
}

// Numeric
func TestForEachIntSlice(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	hof.ForEach(intSlice, func(i int, item int, slice []int) {
		if slice[i] != item {
			t.Errorf("Item is not matching: index=%d, %v", i, item)
		}
		if i < 0 || i >= len(slice) {
			t.Errorf("Index out of range: index=%d, len=%d", i, len(slice))
		}
	})
}

func TestForEachUintSlice(t *testing.T) {
	uintSlice := []uint{1, 2, 3, 4, 5}
	hof.ForEach(uintSlice, func(i int, item uint, slice []uint) {
		if slice[i] != item {
			t.Errorf("Item is not matching: index=%d, %v", i, item)
		}
		if i < 0 || i >= len(slice) {
			t.Errorf("Index out of range: index=%d, len=%d", i, len(slice))
		}
	})
}

func TestForEachUintptrSlice(t *testing.T) {
	uintptrSlice := []uintptr{1, 2, 3, 4, 5}
	hof.ForEach(uintptrSlice, func(i int, item uintptr, slice []uintptr) {
		if slice[i] != item {
			t.Errorf("Item is not matching: index=%d, %v", i, item)
		}
		if i < 0 || i >= len(slice) {
			t.Errorf("Index out of range: index=%d, len=%d", i, len(slice))
		}
	})
}

func TestForEachFloat64Slice(t *testing.T) {
	float64Slice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	hof.ForEach(float64Slice, func(i int, item float64, slice []float64) {
		if slice[i] != item {
			t.Errorf("Item is not matching: index=%d, %v", i, item)
		}
		if i < 0 || i >= len(slice) {
			t.Errorf("Index out of range: index=%d, len=%d", i, len(slice))
		}
	})
}
