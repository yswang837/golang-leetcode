package main

import "fmt"

func main() {
	fmt.Println(QuickSort([]int{4, 2, 3, 4, 5, 65, 45, 5, 6, 67, 6, 3, 5, 43, 17}))
}

func QuickSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start < end {
		povit := partition(nums, start, end)
		quickSort(nums, start, povit-1)
		quickSort(nums, povit+1, end)
	}
}

func partition(nums []int, start, end int) int {
	povitVal := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < povitVal {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, end)
	return i
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
