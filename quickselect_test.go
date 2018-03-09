package main

import (
	"sort"
	"testing"
)

func TestQuickSelect(t *testing.T) {
	arr := []int{2, 1, 9, 7, 3, 5, 5, 6, 8, 1}

	quickSelect(sort.IntSlice(arr), 3)
	expected := []int{2, 1, 1}
	if !hasSameElement(arr[:3], expected) {
		t.Fatalf("Invalid elements got %v expected %v", arr[:3], expected)
	}

	quickSelect(sort.IntSlice(arr), 7)
	expected = []int{2, 1, 1, 3, 5, 5, 6}
	if !hasSameElement(arr[:7], expected) {
		t.Fatalf("Invalid elements got %v expected %v", arr[:7], expected)
	}
}

// hasSameElement chck that two int array contains the same values
func hasSameElement(arr1, arr2 []int) bool {
	m := make(map[int]int)

	if len(arr1) != len(arr2) {
		return false
	}

	for i, elem1 := range arr1 {
		m[elem1]++
		m[arr2[i]]--
	}

	for _, count := range m {
		if count != 0 {
			return false
		}
	}
	return true
}
