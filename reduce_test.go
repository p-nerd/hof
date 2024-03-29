package hof_test

import (
	"testing"

	"github.com/p-nerd/hof"
)

// Unhappy case

func TestReduceNilSlice(t *testing.T) {
	result := hof.Reduce(
		nil,
		func(_ int, acc int, val int) int {
			return acc + val
		},
		0,
	)
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
	result := hof.Reduce(intSlice, nil, 0)
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
	result := hof.Reduce(
		intSlice,
		func(_ int, acc int, val int) int {
			return acc + val
		},
		0,
	)
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
	result := hof.Reduce(
		strSlice,
		func(_ int, acc string, val string) string {
			return acc + val
		},
		"",
	)
	expectedConcat := "Hello World"
	if result != expectedConcat {
		t.Errorf(
			"Result does not match the expected concatenation. Expected: %s, Got: %s",
			expectedConcat,
			result,
		)
	}
}
