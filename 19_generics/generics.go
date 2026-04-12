package main

import "fmt"

// go run 19_generics/generics.go

func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printStringSlice(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// LIFO
type stack[T any] struct {
	elements []T
}

func main() {
	nums := []int{1, 2, 3}
	names := []string{"golang", "typescript"}
	vals := []bool{true, false}

	myStack := stack[string]{
		elements: []string{"golang"},
	}

	// use myStack so it's not declared and unused
	fmt.Println(myStack)

	printStringSlice(names)
	printSlice(vals)
	printSlice(nums)
}
