package stack

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	m := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := []byte{}
	for _, value := range s {
		if value == '(' || value == '[' || value == '{' {
			stack = append(stack, byte(value))
		} else {
			if len(stack) == 0 { // 防止这种括号 }{
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if m[byte(value)] != top {
				return false
			}
		}
	}
	return len(stack) == 0
}
