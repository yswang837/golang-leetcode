package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 0: hot-100 内的任意题目
// 1: 哈希
// 2: 双指针
// 3: 滑动窗口
// 4: 子串
// 5: 普通数组
// 6: 矩阵
// 7: 链表
// 8: 二叉树
// 9: 图论
// 10: 回溯
// 11: 二分查找
// 12: 栈
// 13: 堆
// 14: 贪心算法
// 15: 动态规划
// 16: 多维动态规划
// 17: 技巧

type problem struct {
	t       int
	numbers []string
}

type problems struct {
	kind []problem
}

func newProblems() *problems {
	return &problems{[]problem{
		{1, []string{"1.两数之和", "49.字母异位词分组", "128.最长连续序列"}},
		{2, []string{"283.", "11.", "15.", "42."}},
		{3, []string{"3.", "438."}},
		{4, []string{"560.", "239.", "76."}},
		{5, []string{"53.", "56.", "189.", "238.", "41."}},
		{6, []string{"73.", "54.", "48.", "240."}},
		{7, []string{"160.", "206.", "146.", "234.", "141.", "142.", "21.", "2.", "19.", "24.", "25.", "138.", "148.", "23."}},
		{8, []string{"94.", "104.", "226.", "101.", "543.", "102.", "108.", "98.", "230.", "199.", "114.", "105.", "437.", "236.", "124."}},
		{9, []string{"200.", "994.", "207.", "208."}},
		{10, []string{"46.", "78.", "17.", "39.", "22.", "79.", "131.", "51."}},
		{11, []string{"35.", "74.", "34.", "33.", "153.", "4."}},
		{12, []string{"20.", "155.", "394.", "739.", "84."}},
		{13, []string{"215.", "347.", "295."}},
		{14, []string{"121.", "55.", "45.", "763."}},
		{15, []string{"70.", "118.", "198.", "279.", "322.", "139.", "300.", "152.", "416.", "32."}},
		{16, []string{"62.", "64.", "5.", "1143.", "72."}},
		{17, []string{"136.", "169.", "75.", "31.", "287."}},
	}}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	p := newProblems()
	// 0: hot-100 内的任意题目
	// 1: 哈希
	// 2: 双指针
	// 3: 滑动窗口
	// 4: 子串
	// 5: 普通数组
	// 6: 矩阵
	// 7: 链表
	// 8: 二叉树
	// 9: 图论
	// 10: 回溯
	// 11: 二分查找
	// 12: 栈
	// 13: 堆
	// 14: 贪心算法
	// 15: 动态规划
	// 16: 多维动态规划
	// 17: 技巧
	fmt.Printf("今天做 %s ....\n", p.randProblem(8))
}

func (p *problems) randProblem(t int) string {
	if t < 0 || t > 17 {
		return ""
	}
	if t == 0 {
		allProblem := []string{}
		for _, val := range p.kind {
			allProblem = append(allProblem, val.numbers...)
		}
		return allProblem[rand.Intn(100)]
	} else {
		numbers := p.kind[t-1].numbers
		return numbers[rand.Intn(len(numbers))]
	}

}
