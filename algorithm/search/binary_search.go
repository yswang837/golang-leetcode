package main

import (
	"fmt"
)

func main() {
	fmt.Println(binarySearch([]int{2, 3, 3, 4, 4, 5, 5, 5, 6, 6, 17, 43, 45, 65, 67}, 65))
}

// Search 二分搜索最常用模板
func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, middle, right := 0, 0, len(nums)-1
	for left <= right {
		middle = (left + right) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle
		} else {
			left = middle
		}
	}
	return -2
}
