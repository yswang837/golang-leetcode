package binary_search

import "fmt"

func search(nums []int, target int) int {
	//将数组一分为二，其中一定有一个是有序的，另一个可能是有序，也能是部分有序。
	// 此时有序部分用二分法查找。无序部分再一分为二，其中一个一定有序，另一个可能有序，可能无序。就这样循环。
	//判断target在不在有序数组里，只需要比较最小值和最大值即可，如果target不在顺序区间，那么便去乱序区间，在其中再分割出有序区间。如果有序区间中有target那便好说，直接二分便可获得答案
	left, middle, right := 0, 0, len(nums)-1
	for left <= right {
		middle = (left + right) / 2
		if nums[middle] == target {
			return middle
		}
		if nums[left] <= nums[middle] {
			if target < nums[middle] && target >= nums[left] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else {
			if target > nums[middle] && target <= nums[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 1, 2, 3}, 3))
}
