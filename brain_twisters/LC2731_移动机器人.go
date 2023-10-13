package brain_twisters

import "sort"

// 碰撞后朝相反方向运动,相当于没有碰撞
func sumDistance(nums []int, s string, d int) int {
	const mod int = 1e9 + 7
	// 计算时间d之后机器人的位置
	for i := range nums {
		if s[i] == 'L' {
			nums[i] -= d
		} else {
			nums[i] += d
		}
	}
	// 对机器人坐标排序
	sort.Ints(nums)
	// 求和,遍历所有机器人,第i个机器人与在它前面的所有机器人的距离求和
	sum, res := 0, 0
	for i := 1; i < len(nums); i++ {
		sum += nums[i-1]
		res = (res + i*nums[i] - sum) % mod
	}
	return res
}
