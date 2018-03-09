package main

import (
	"math/rand"
	"sort"
)

// quickSelect updates in place a sort.Interface such as the first k element are the k smallest
func quickSelect(data sort.Interface, k int) {
	if k < 1 || k > data.Len() {
		// No sorting to do as it's already done
		return
	}

	pivotSelect(data, k, 0, data.Len()-1)
}

func pivotSelect(data sort.Interface, k, min, max int) {
	var pivot int

	for {
		if min >= max {
			return
		}

		pivot = rand.Intn(max+1-min) + min
		pivot = partition(data, pivot, min, max)

		if k == pivot {
			return
		} else if k < pivot {
			max = pivot - 1
		} else {
			min = pivot + 1
		}
	}
}

func partition(data sort.Interface, pivot, min, max int) int {
	// Put the pivot in the right most position
	data.Swap(pivot, max)

	// Index of the leftmost item seen whose isn't compared to pivot
	storeIndex := min
	for i := min; i < max; i++ {
		// If the item is smaller than pivot move it in place of the storeIndex
		if data.Less(i, max) {
			data.Swap(i, storeIndex)
			storeIndex++
		}
	}
	data.Swap(storeIndex, max)
	return storeIndex
}
