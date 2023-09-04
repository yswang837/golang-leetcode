//https://leetcode.cn/problems/group-anagrams/description/?envType=study-plan-v2&envId=top-100-liked

func groupAnagrams(strs []string) [][]string {
	ret := [][]string{}
	if len(strs) == 0 {
		return ret
	}
	if len(strs) == 1 {
		ret = append(ret,strs)
		return ret
	}
	m := map[string][]string{}
	var b []byte
	for _,str := range strs {
		b = []byte(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		key := string(b)
		itemSlice := append(m[key],str)
		m[key] = itemSlice
	}
	for _,val := range m {
		ret = append(ret,val)
	}
	return ret
}
