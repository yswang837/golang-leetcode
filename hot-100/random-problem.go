package main

import (
	"fmt"
	"math/rand"
	"time"
)

type problem struct {
	t       int
	numbers []string
}

type problems struct {
	kind []problem
}

func newProblems() *problems {
	return &problems{[]problem{
		{1, []string{"1. 两数之和", "49. 字母异位词分组", "128. 最长连续序列"}},
		{2, []string{"283. 移动零", "11. 盛最多水的容器", "15. 三数之和", "42. 接雨水"}},
		{3, []string{"3. 无重复字符的最长子串", "438. 找到字符串中所有字母异位词"}},
		{4, []string{"560. 和为K的字数组", "239. 滑动窗口最大值", "76. 最小覆盖子串"}},
		{5, []string{"53. 最大子数组和", "56. 合并区间", "189. 轮转数组", "238. 除自身以外的数组的乘积", "41. 缺失的第一个正数"}},
		{6, []string{"73. 矩阵置零", "54. 旋转矩阵", "48. 旋转图像", "240. 搜索二维矩阵II"}},
		{7, []string{"160. 相交链表", "206. 反转链表", "146. LRU缓存", "234. 回文链表", "141. 环形链表", "142. 环形链表II", "21. 合并两个有序链表", "2. 两数相加", "19. 删除链表的倒数第N个结点", "24. 两两交换链表中的节点", "25. K个一组翻转链表", "138. 随机链表的复制", "148. 排序链表", "23. 合并K个升序链表"}},
		{8, []string{"94. 二叉树的中序遍历", "104. 二叉树的最大深度", "226. 翻转二叉树", "101. 对称二叉树", "543. 二叉树的直径", "102. 二叉树的层序遍历", "108. 将有序数组转换为二叉搜索树", "98. 验证二叉搜索树", "230. 二叉搜索树中第K小的元素", "199. 二叉树的右视图", "114. 二叉树展开为链表", "105. 从前序与中序遍历序列构造二叉树", "437. 路径总和III", "236. 二叉树的最近公共祖先", "124. 二叉树中的最大路径和"}},
		{9, []string{"200. 岛屿数量", "994. 腐烂的橘子", "207. 课程表", "208. 实现Trie(前缀树)"}},
		{10, []string{"46. 全排列", "78. 子集", "17. 电话号码的字母组合", "39. 组合总和", "22. 括号生成", "79. 单词搜索", "131. 分割回文串", "51. N皇后"}},
		{11, []string{"35. 搜索插入位置", "74. 搜索二维矩阵", "34. 在排序数组中查找元素的第一个和最后一个位置", "33. 搜索旋转排序数组", "153. 寻找旋转排序数组中的最小值", "4. 寻找两个正序数组的中位数"}},
		{12, []string{"20. 有效的括号", "155. 最小栈", "394. 字符串解码", "739. 每日温度", "84. 柱状图中最大的矩形"}},
		{13, []string{"215. 数组中的第K个最大元素", "347. 前K个高频元素", "295. 数据流的中位数"}},
		{14, []string{"121. 买卖股票的最佳时机", "55. 跳跃游戏", "45. 跳跃游戏II", "763. 划分字母区间"}},
		{15, []string{"70. 爬楼梯", "118. 杨辉三角", "198. 打家劫舍", "279. 完全平方数", "322. 零钱兑换", "139. 单词拆分", "300. 最长递增子序列", "152. 乘积最大子数组", "416. 分割等和子集", "32. 最长有效括号"}},
		{16, []string{"62. 不同路径", "64. 最小路径和", "5. 最长回文子串", "1143. 最长公共子序列", "72. 编辑距离"}},
		{17, []string{"136. 只出现一次的数字", "169. 多数元素", "75. 颜色分类", "31. 下一个排列", "287. 寻找重复数"}},
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
	fmt.Printf("今天做 %s\n", p.randProblem(8))
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
