package main

import (
	"fmt"
	"sort"
)

func ThreeNumberSum(array []int, targetSum int) [][]int {
	sort.Ints(array)
	triplets := make([][]int, 0)

	for i := 0; i < len(array)-2; i++ {
		leftIdx := i + 1
		rightIdx := len(array) - 1

		for leftIdx < rightIdx {
			currentNum := array[i]
			leftNum := array[leftIdx]
			rightNum := array[rightIdx]
			potentialSum := leftNum + rightNum + currentNum

			if targetSum == potentialSum {
				triplets = append(triplets, []int{currentNum, leftNum, rightNum})
				rightIdx--
				leftIdx++
			} else if potentialSum > targetSum {
				rightIdx--
			} else if potentialSum < targetSum {
				leftIdx++
			}
		}
	}

	return triplets

}

func main() {
	fmt.Println(ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, 0, -8, -1}, 0))
	fmt.Println(ThreeNumberSum([]int{8, 10, -2, 49, 14}, 57))
	fmt.Println(ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, 0, -8, -1, 6, -5}, 0))
	fmt.Println(ThreeNumberSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}, 18))
}
