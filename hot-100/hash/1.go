//https://leetcode.cn/problems/two-sum/?envType=study-plan-v2&envId=top-100-liked

// 解题要点：一次遍历，把num作为map的key，index作为value，可高效解题

package main

func twoSum(nums []int, target int) []int {
	ret := []int{-1, -1}
	if len(nums) == 0 {
		return ret
	}
	m := map[int]int{}
	for i, val := range nums {
		tmp := target - val
		if i2, ok := m[tmp]; ok {
			return []int{i, i2}
		}
		m[val] = i
	}
	return ret
}
