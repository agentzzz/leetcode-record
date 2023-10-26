package string

import "strconv"

func countSeniors(details []string) int {
	num := 0
	for _, detail := range details {
		age := getAgeFromString(detail)
		if age > 60 {
			num++
		}
	}
	return num
}

func getAgeFromString(s string) int {
	if len(s) < 13 {
		return 0
	}
	ageStr := s[11:13]
	age, _ := strconv.ParseInt(ageStr, 10, 32)
	return int(age)
}
