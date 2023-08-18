package main

import "fmt"

func main() {
	fmt.Println(Search([]int{2, 3, 3, 4, 4, 5, 5, 5, 6, 6, 17, 43, 45, 65, 67}, 17))
}

// Search 二分搜索最常用模板
func Search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high + low) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		}
	}
	return -1
}
