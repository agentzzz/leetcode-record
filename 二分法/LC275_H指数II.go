package 二分法

func hIndex(citations []int) int {
	i, j := 0, len(citations)-1
	for i <= j {
		mid := i + (j-i)>>1
		if citations[mid] >= len(citations)-mid {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return len(citations) - i
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
