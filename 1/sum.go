// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2020-12-18 9:49
// version: 1.0.0
// desc   :

package sum

// 暴力枚举
// 遍历两个数组，每次遍历都计算一次 x + y == target
func _1(arr []int, target int) (res []int) {
	for i, x := range arr {
		for j, y := range arr {
			if x+y == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 哈希表
func _2(arr []int, target int) (res []int) {
	mp := make(map[int]int)
	for i, x := range arr {
		if p, ok := mp[target-x]; ok {
			return []int{p, i}
		}
		mp[x] = i
	}
	return nil
}
