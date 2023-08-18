package main

import "fmt"

func main() {
	fmt.Println(Search([]int{2, 3, 3, 4, 4, 5, 5, 5, 6, 6, 17, 43, 45, 65, 67}, 4))
}

// Search 二分搜索最常用模板
func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	for low <= high {
		middle := (low + high) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return -2
}
