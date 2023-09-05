package main

import (
	"fmt"
	"math"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	hashMap := make(map[int]bool)
	for _, num := range nums {
		hashMap[num] = true
	}

	// 3.遍历每个数，执行并查操作
	res := math.MinInt
	for num, _ := range hashMap {
		// 3.1 查：查一组连续序列的第一个数（即：判断该数前面还有没有数）
		// 如果当前数不是一组连续序列的第一个数，就跳过
		// 如果当前数是一组连续序列的第一个数，就寻找以该数开头的后续序列
		if ok := hashMap[num-1]; ok {
			continue
		}
		flag := num        // 表示当前找到了这组连续序列中的哪个数，初始化为这组连续序列的第一个数num
		currentLength := 1 // 初始化以该num开头的连续序列长度为1
		// 3.2 并：在找到一组连续序列的第一个数后，就寻找以该数开头的后续序列，将后续序列归并，直到找到这组序列的最后一个数，计算这组序列的长度
		for {
			if ok := hashMap[flag+1]; ok {
				// 如果当前数的下一个数在数组中，就归并，flag+1更新当前找到连续序列中的哪个数，currentLength+1更新这组序列长度
				flag++
				currentLength++
			} else {
				// 如果当前数的下一个数不在数组中，说明该组连续序列已经找完了，此时跳出循环，寻找下一个连续序列
				break
			}
		}
		// 3.3 找到每组序列的最长的那一个序列，就是我们要找的结果集
		res = maxInt(res, currentLength)
	}

	// 4.返回结果集
	return res
}

// 求两个int数的较大值
func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
