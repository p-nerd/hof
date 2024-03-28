package hof_test

import (
	"testing"

	"github.com/p-nerd/hof"
)

// Unhappy case

func TestReduceNilSlice(t *testing.T) {
	result, err := hof.Reduce(
		nil,
		func(_ int, acc int, val int) int {
			return acc + val
		},
		0,
	)
	if err == nil {
		t.Error("Expected error for nil slice, got nil")
	} else if err.Error() != "slice is nil" {
		t.Errorf("Expected error message 'slice is nil', got '%s'", err.Error())
	}
	expectedInitial := 0
	if result != expectedInitial {
		t.Errorf(
			"Expected result to match the initial value for nil slice input. Expected: %d, Got: %d",
			expectedInitial,
			result,
		)
	}
}

func TestReduceNilReducer(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	result, err := hof.Reduce(intSlice, nil, 0)
	if err == nil {
		t.Error("Expected error for nil reducer function, got nil")
	} else if err.Error() != "reducer is nil" {
		t.Errorf("Expected error message 'reducer is nil', got '%s'", err.Error())
	}
	expectedInitial := 0
	if result != expectedInitial {
		t.Errorf(
			"Expected result to match the initial value for nil reducer function. Expected: %d, Got: %d",
			expectedInitial,
			result,
		)
	}
}

// Happy case

func TestReduceSumInts(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	result, err := hof.Reduce(
		intSlice,
		func(_ int, acc int, val int) int {
			return acc + val
		},
		0,
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedSum := 15
	if result != expectedSum {
		t.Errorf(
			"Result does not match the expected sum. Expected: %d, Got: %d",
			expectedSum,
			result,
		)
	}
}

func TestReduceConcatStrings(t *testing.T) {
	strSlice := []string{"Hello", " ", "World"}
	result, err := hof.Reduce(
		strSlice,
		func(_ int, acc string, val string) string {
			return acc + val
		},
		"",
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedConcat := "Hello World"
	if result != expectedConcat {
		t.Errorf(
			"Result does not match the expected concatenation. Expected: %s, Got: %s",
			expectedConcat,
			result,
		)
	}
}
