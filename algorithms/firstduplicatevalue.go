package main

import "fmt"

func FirstDuplicateValue(array []int) int {
	arraysMap := make(map[int]int)

	for _, num := range array {
		if _, ok := arraysMap[num]; ok {
			arraysMap[num]++
		} else {
			arraysMap[num] = 0
		}

		if arraysMap[num] > 1 {
			return num
		}
	}

	return -1
}

func main() {
	fmt.Println(FirstDuplicateValue([]int{2, 1, 5, 2, 3, 3, 4}))
	fmt.Println(FirstDuplicateValue([]int{2, 1, 5, 3, 3, 2, 4}))
	fmt.Println(FirstDuplicateValue([]int{}))
	fmt.Println(FirstDuplicateValue([]int{1}))
	fmt.Println(FirstDuplicateValue([]int{1, 1}))
}
