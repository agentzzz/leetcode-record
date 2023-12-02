package JZ_Offer

func Find(target int, array [][]int) bool {
	if len(array) == 0 || len(array[0]) == 0 {
		return false
	}
	n := len(array[0])
	i, j := len(array)-1, 0
	for i >= 0 && j < n {
		if array[i][j] == target {
			return true
		} else if array[i][j] < target {
			j++
		} else if array[i][j] > target {
			i--
		}
	}
	return false
}
