package hof_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/p-nerd/hof"
)

func TestMapIntSlice(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4, 6, 8, 10}
	result, err := hof.Map(input, func(_ int, item int, _ []int) int { return item * 2 })
	if err != nil {
		t.Errorf("Error occurred during map operation: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test for integers failed: expected %v, got %v", expected, result)
	}
}

func TestMapStringSlice(t *testing.T) {
	input := []string{"hello", "world"}
	expected := []string{"HELLO", "WORLD"}
	result, err := hof.Map(
		input,
		func(_ int, item string, _ []string) string { return strings.ToUpper(item) },
	)
	if err != nil {
		t.Errorf("Error occurred during map operation: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test for integers failed: expected %v, got %v", expected, result)
	}
}

type User struct {
	ID   int
	Name string
}

func TestMapStruct(t *testing.T) {
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}
	expectedIDs := []int{1, 2, 3}
	resultIDs, err := hof.Map(users, func(_ int, user User, _ []User) int { return user.ID })
	if err != nil {
		t.Errorf("Error occurred during map operation: %v", err)
	}
	if !reflect.DeepEqual(resultIDs, expectedIDs) {
		t.Errorf("Test for structs failed: expected %v, got %v", expectedIDs, resultIDs)
	}
}
