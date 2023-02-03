package main

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
