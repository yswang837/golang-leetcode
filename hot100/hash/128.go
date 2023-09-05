//https://leetcode.cn/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-100-liked

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	hashMap := make(map[int]bool)
	for _, num := range nums {
		hashMap[num] = true
	}

	// 3.遍历每个数，执行并查操作
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
