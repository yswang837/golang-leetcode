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

// 原地交换，所以传入交换索引
func quickSort(nums []int, start, end int) {
	if start < end {
		// 分治法：divide
		pivot := partition(nums, start, end)
		quickSort(nums, 0, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

// 分区
func partition(nums []int, start, end int) int {
	// 选取最后一个元素作为基准pivot
	p := nums[end]
	i := start
	// 最后一个值就是基准所以不用比较
	for j := start; j < end; j++ {
		if nums[j] < p {
			swap(nums, i, j)
			i++
		}
	}
	// 把基准值换到中间
	swap(nums, i, end)
	return i
}

// 交换两个元素
func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
