package main

import "fmt"

func main() {
	fmt.Println(Search([]int{2, 3, 3, 4, 4, 5, 5, 5, 6, 6, 17, 43, 45, 65, 67}, 17))
}

// Search 二分搜索最常用模板
func Search(nums []int, target int) int {
	// 1、初始化start、end
	start := 0
	end := len(nums) - 1
	// 2、处理for循环
	for start+1 < end {
		mid := start + (end-start)/2
		// 3、比较a[mid]和target值
		if nums[mid] == target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		} else if nums[mid] > target {
			end = mid
		}
	}
	// 4、最后剩下两个元素，手动判断
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}
