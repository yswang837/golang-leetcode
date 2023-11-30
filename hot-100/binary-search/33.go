package binary_search

import "fmt"

// 将数组一分为二，其中一定有一个是有序的，另一个可能是有序，也能是部分有序。
// 此时有序部分用二分法查找。无序部分再一分为二，其中一个一定有序，另一个可能有序，可能无序。就这样循环。
// 判断target在不在有序数组里，只需要比较最小值和最大值即可，如果target不在顺序区间，那么便去乱序区间，在其中再分割出有序区间。如果有序区间中有target那便好说，直接二分便可获得答案
// [4, 5, 6, 7, 8, 9, 1, 2, 3], 3  第一次循环左边有序
// [7, 8, 9, 1, 2, 3, 4, 5, 6], 3  第一次循环右边有序
func search(nums []int, target int) int {
	left, middle, right := 0, 0, len(nums)-1
	for left <= right {
		middle = (left + right) / 2
		if nums[middle] == target {
			return middle
		}
		// 判断哪边有序
		if nums[left] <= nums[middle] {
			// 左边有序
			if nums[left] <= target && target < nums[middle] {
				// 目标值在这个有序区间
				right = middle - 1
			} else {
				// 目标值在这个无序区间
				left = middle + 1
			}
		} else {
			// 右边有序
			if nums[middle] < target && target <= nums[right] {
				// 目标值在这个有序区间
				left = middle + 1
			} else {
				// 目标值在这个无序区间
				right = middle - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 1, 2, 3}, 3))
}
