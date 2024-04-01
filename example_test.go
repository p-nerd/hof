package hof_test

import (
	"fmt"

	"github.com/p-nerd/hof"
)

func ExampleForEach() {
	hof.ForEach(
		[]string{"apple", "banana", "cherry"},
		func(index int, item string, _ []string) {
			fmt.Println(index, item)
		},
	)
	// Output:
	// 0 apple
	// 1 banana
	// 2 cherry
}

func ExampleMap() {
	mappedSlice := hof.Map([]int{1, 2, 3, 4, 5}, func(index int, item int, slice []int) int {
		return item * 2
	})
	fmt.Println(mappedSlice)
	// Output: [2 4 6 8 10]
}

func ExampleFilter() {
	filteredSlice := hof.Filter(
		[]int{1, 2, 3, 4, 5},
		func(index int, item int, slice []int) bool {
			return item%2 != 0
		},
	)
	fmt.Println(filteredSlice)
	// Output: [1 3 5]
}

func ExampleReduce() {
	sum := hof.Reduce([]int{1, 2, 3, 4, 5}, func(index int, acc int, val int) int {
		return acc + val
	}, 0)
	fmt.Println(sum)
	// Output: 15
}
