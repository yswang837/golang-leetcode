package double_point

func maxArea(height []int) int {
	if len(height) == 0 || len(height) == 1 {
		return 0
	}
	left, right := 0, len(height)-1
	maxarea := 0
	for left < right {
		h, flag := subabs(height[left], height[right])
		w := right - left
		maxarea = maxInt(h*w, maxarea)
		if flag {
			right--
		} else {
			left++
		}
	}
	return maxarea
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func subabs(i, j int) (int, bool) {
	//左是否大于右
	if i-j < 0 {
		return i, false
	}
	return j, true
}
