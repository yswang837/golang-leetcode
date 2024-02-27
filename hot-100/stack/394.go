package stack

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	numStack := []int{}
	strStack := []string{}
	num := 0  // 用于累积读取完整数字
	ret := "" // 用于累积当前解码的字符串
	for _, char := range s {
		switch char {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			n, _ := strconv.Atoi(string(char))
			num = num*10 + n
		case '[': // 当遇到左括号，表示一个新的编码字符串开始
			strStack = append(strStack, ret)
			ret = "" // 重置数字和字符串，以便处理新的编码字符串
			numStack = append(numStack, num)
			num = 0
		case ']': // 当遇到右括号，表示当前编码字符串结束
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			ret = str + strings.Repeat(ret, count)
		default: // 当遇到字母时，直接添加到当前字符串中
			ret += string(char)
		}
	}
	return ret
}
