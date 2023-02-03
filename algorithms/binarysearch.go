package main

import "testing"

func BinarySearch(array []int, target int) int {
	leftIdx := 0
	rightIdx := len(array)

	for leftIdx <= rightIdx {
		mid := (leftIdx + rightIdx) / 2
		potentialTarget := array[mid]

		if potentialTarget == target {
			return potentialTarget
		} else if potentialTarget > target {
			rightIdx = mid - 1
		} else if potentialTarget < target {
			leftIdx = mid + 1
		}
	}
	return -1
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		array  []int
		target int
		result int
	}{
		{[]int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}, 33, 3},
		{[]int{1, 5, 23, 111}, 111, 3},
		{[]int{1, 5, 23, 111}, 5, 1},
		{[]int{1, 5, 23, 111}, 35, -1},
		{[]int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}, 0, 0},
	}
	for _, test := range tests {
		if got := BinarySearch(test.array, test.target); got != test.result {
			t.Errorf("binarySearch(%v, %d) = %d; expected %d", test.array, test.target, got, test.result)
		}
	}
}
