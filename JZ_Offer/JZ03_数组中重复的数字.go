package JZ_Offer

func duplicate(numbers []int) int {
	memo := make(map[int]bool)
	for _, num := range numbers {
		if memo[num] {
			return num
		}
		memo[num] = true
	}
	return -1
}

// 第二种方法还有点问题,需要调试一下
func duplicate2(numbers []int) int {
	for i, num := range numbers {
		if num == numbers[num] {
			return num
		}
		swap(numbers, i, num)
	}
	return -1
}

func swap(numbers []int, i, j int) {
	temp := numbers[i]
	numbers[i] = numbers[j]
	numbers[j] = temp
}
