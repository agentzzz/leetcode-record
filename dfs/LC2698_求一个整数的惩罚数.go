package dfs

import "strconv"

func punishmentNumber(n int) int {
	var check func(string, int, int, int) bool
	check = func(str string, start, pathSum, target int) bool {
		if start == len(str) {
			return pathSum == target
		}
		temp := 0
		for i := start; i < len(str); i++ {
			temp = temp*10 + int(str[i]-'0')
			if pathSum+temp > target {
				break
			}
			if check(str, start+1, pathSum+temp, target) {
				return true
			}
		}
		return false
	}
	res := 0
	for i := 1; i <= n; i++ {
		if check(strconv.Itoa(i*i), 0, 0, i) {
			res += i * i
		}
	}
	return res
}
