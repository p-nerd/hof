package hof_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/p-nerd/hof"
)

func TestForEach(t *testing.T) {
	tt := []struct {
		name   string
		called bool
		stype  string
		slice  []any
	}{
		{"nil slice", false, "", []any{}},
		{"empty slice", false, "", []any{}},
		{"valid value", true, "int", []any{1}},
		{"boolean slice", true, "bool", []any{true, false, true}},
		{"string slice", true, "string", []any{"apple", "banana", "cherry"}},
		{"int slice", true, "int", []any{1, 2, 3, -4, 5, -1}},
		{"uint slice", true, "uint", []any{uint(1), uint(3), uint(5)}},
		{"float64 slice", true, "float64", []any{1.1, 2.2, 3.3, 4.4, 5.5}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			hof.ForEach(tc.slice, func(index int, item any, slice []any) {
				if !tc.called {
					t.Errorf("Callback should not be call but it called")
					return
				}
				if stype := fmt.Sprintf("%T", item); stype != tc.stype {
					t.Errorf("Item type should be %v; got %v", tc.stype, stype)
					return
				}
				if slice[index] != item {
					t.Errorf("Item is not matching: index=%d, %v", index, item)
				}
			})
		})
	}

	t.Run("nil callback", func(t *testing.T) {
		hof.ForEach([]int{1, 2, 3}, nil)
	})
}

func TestMap(t *testing.T) {
	tt := []struct {
		name string
		test func(*testing.T, string)
	}{
		{
			"nil slice",
			func(t *testing.T, name string) {
				testCaseMap(
					t,
					name,
					nil,
					[]int{},
					func(_ int, item int, _ []int) int { return item },
				)
			},
		},
		{
			"nil callback",
			func(t *testing.T, name string) {
				testCaseMap(
					t,
					name,
					[]int{},
					[]int{},
					nil,
				)
			},
		},
		{
			"int slice double",
			func(t *testing.T, name string) {
				testCaseMap(
					t,
					name,
					[]int{1, 2, 3, 4, 5},
					[]int{2, 4, 6, 8, 10},
					func(_ int, item int, _ []int) int { return item * 2 },
				)
			},
		},
		{
			"string slice upper case",
			func(t *testing.T, name string) {
				testCaseMap(
					t,
					name,
					[]string{"hello", "world"},
					[]string{"HELLO", "WORLD"},
					func(_ int, item string, _ []string) string { return strings.ToUpper(item) },
				)
			},
		},
		{
			"user slice return one key",
			func(t *testing.T, name string) {
				type User struct {
					ID   int
					Name string
				}
				testCaseMap(
					t,
					name,
					[]User{
						{ID: 1, Name: "Alice"},
						{ID: 2, Name: "Bob"},
						{ID: 3, Name: "Charlie"},
					},
					[]int{1, 2, 3},
					func(_ int, user User, _ []User) int { return user.ID },
				)
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.test(t, tc.name)
		})
	}
}

func testCaseMap[T, U any](
	t *testing.T,
	name string,
	input []T,
	expected []U,
	callback func(int, T, []T) U,
) {
	result := hof.Map(input, callback)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Testing %v: expected %v, got %v", name, expected, result)
	}
}

func TestFilter(t *testing.T) {
	testCases := []struct {
		name          string
		inputSlice    []int
		filterFunc    func(int, int, []int) bool
		expectedSlice []int
	}{
		{
			name:          "Nil slice",
			inputSlice:    nil,
			filterFunc:    func(_ int, item int, _ []int) bool { return true },
			expectedSlice: []int{},
		},
		{
			name:          "Nil callback",
			inputSlice:    []int{1, 2, 3, 4, 5},
			filterFunc:    nil,
			expectedSlice: []int{},
		},
		{
			name:          "Filter even numbers",
			inputSlice:    []int{1, 2, 3, 4, 5},
			filterFunc:    func(_ int, item int, _ []int) bool { return item%2 == 0 },
			expectedSlice: []int{2, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			filteredSlice := hof.Filter(tc.inputSlice, tc.filterFunc)
			if !reflect.DeepEqual(filteredSlice, tc.expectedSlice) {
				t.Errorf(
					"Filtered slice does not match the expected result. Expected: %v, Got: %v",
					tc.expectedSlice,
					filteredSlice,
				)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tt := []struct {
		name string
		test func(*testing.T, string)
	}{
		{
			"nil slice",
			func(t *testing.T, name string) {
				testCaseReduce(
					t,
					name,
					nil,
					func(int, int, int) int { return 0 },
					0,
					0,
				)
			},
		},
		{
			"nil reducer",
			func(t *testing.T, name string) {
				testCaseReduce(
					t,
					name,
					[]int{1, 2, 3, 4, 5},
					nil,
					0,
					0,
				)
			},
		},
		{
			"sum integers",
			func(t *testing.T, name string) {
				testCaseReduce(
					t,
					name,
					[]int{1, 2, 3, 4, 5},
					func(_ int, acc int, val int) int { return acc + val },
					0,
					15,
				)
			},
		},
		{
			"concatenate strings",
			func(t *testing.T, name string) {
				testCaseReduce(
					t,
					name,
					[]string{"Hello", " ", "World"},
					func(_ int, acc string, val string) string { return acc + val },
					"",
					"Hello World",
				)
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.test(t, tc.name)
		})
	}
}

func testCaseReduce[T, U any](
	t *testing.T,
	name string,
	input []T,
	reducer func(int, U, T) U,
	initialValue U,
	expectedResult U,
) {
	result := hof.Reduce(input, reducer, initialValue)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Testing %v: expected %v, got %v", name, expectedResult, result)
	}
}
