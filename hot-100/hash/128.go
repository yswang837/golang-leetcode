//https://leetcode.cn/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-100-liked

// 解题要点：将切片放入map，找到每段的起始位置，从map里面取每段的值，两两求max就行。

package main

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	hashMap := make(map[int]bool)
	for _, num := range nums {
		hashMap[num] = true
	}
	res := 0
	for num, _ := range hashMap {
		if ok := hashMap[num-1]; ok {
			continue
		}
		flag := num
		currentLength := 1
		for {
			if ok := hashMap[flag+1]; ok {
				flag++
				currentLength++
			} else {
				break
			}
		}
		res = maxInt(res, currentLength)
	}
	return res
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
