package dfs

//func punishmentNumber(n int) int {
//	sum := 0
//	for i := 1; i <= n; i++ {
//		target := i * i
//		targetStr := strconv.Itoa(target)
//		if check(targetStr, target, 0, 0) {
//			sum += target
//		}
//	}
//	return sum
//}
//
//func check(targetStr string, target, sum, start int) bool {
//	// 出口
//	if start == len(targetStr) {
//		return target == sum
//	}
//	// 分支
//	for i := start; i < len(targetStr); i++ {
//		temp := 0
//		temp = temp*10 + int(targetStr[i]-'0')
//		if sum+temp > target {
//			break
//		}
//		if check(targetStr, target, sum+temp, start+1) {
//			return true
//		}
//	}
//	return false
//}
