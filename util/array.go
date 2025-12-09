package util

import "strconv"

func StrToIntArr(s []string) []int {
	arr := make([]int, len(s))

	for i := range s {
		intValue, err := strconv.Atoi(s[i])
		if err != nil {
			panic("invalid value to convert")
		}
		arr[i] = intValue
	}
	return arr
}
