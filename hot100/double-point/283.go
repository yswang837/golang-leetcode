package double_point

//https://leetcode.cn/problems/move-zeroes/?envType=study-plan-v2&envId=top-100-liked

func moveZeroes(nums []int) {
	slow, fast, length := 0, 0, len(nums)
	for fast < length {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}

//           f
//      s
// [0,1,0,3,12]
// [1,0,0,3,12]
// [1,3,0,0,12]
// [1,3,12,0,0]
