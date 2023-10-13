package brain_twisters

import (
	"fmt"
	"strconv"
	"testing"
)

func findTheArrayConcVal(nums []int) int64 {
	var num int64
	// 双指针
	i, j := 0, len(nums)-1
	for i <= j {
		if i == j {
			num += int64(nums[i])

		} else {
			num += AddFirstAndLast(nums[i], nums[j])
		}
		i++
		j--
	}
	return num
}

func AddFirstAndLast(first, end int) int64 {
	firstStr, endStr := strconv.Itoa(first), strconv.Itoa(end)
	resultStr := firstStr + endStr
	result, _ := strconv.ParseInt(resultStr, 10, 64)
	return result
}

func TestArray(t *testing.T) {
	nums := []int{5, 14, 13, 8, 12}
	res := findTheArrayConcVal(nums)
	fmt.Println(res)
}
