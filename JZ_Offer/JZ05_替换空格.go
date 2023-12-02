package JZ_Offer

func replaceSpace(s string) string {
	// write code here
	res := make([]rune, 0)
	c := []rune("%20")
	for _, v := range s {
		if v == ' ' {
			res = append(res, c...)
		} else {
			res = append(res, v)
		}
	}
	return string(res)
}
