package string

import "strconv"

func countDigits(num int) int {
	numStr := strconv.Itoa(num)
	res := 0
	for _, item := range numStr {
		val := int(item - '0')
		if num%val == 0 {
			res++
		}
	}
	return res
}
