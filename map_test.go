package hof_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/p-nerd/hof"
)

func TestMapInt(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4, 6, 8, 10}
	result := hof.Map(input, func(x int) int { return x * 2 })
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test for integers failed: expected %v, got %v", expected, result)
	}
}

func TestMapString(t *testing.T) {
	inputStr := []string{"hello", "world"}
	expectedStr := []string{"HELLO", "WORLD"}
	resultStr := hof.Map(inputStr, func(s string) string { return strings.ToUpper(s) })
	if !reflect.DeepEqual(resultStr, expectedStr) {
		t.Errorf("Test for strings failed: expected %v, got %v", expectedStr, resultStr)
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
	resultIDs := hof.Map(users, func(user User) int { return user.ID })
	if !reflect.DeepEqual(resultIDs, expectedIDs) {
		t.Errorf("Test for structs failed: expected %v, got %v", expectedIDs, resultIDs)
	}
}
