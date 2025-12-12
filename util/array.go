package util

import "strconv"

func StrArrToIntArr(s []string) []int {
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

func StrToIntArr(s string) []int {
	var ints []int
	buffer := Number(0)
	set := false
	neg := false
	for _, char := range s {
		if IsNum(char) {
			buffer.Add(NumVal(char))
			set = true
			continue
		}
		if char == '-' {
			if set {
				b := buffer.Clear()
				if neg {
					b = -b
				}
				ints = append(ints, b)
				set = false
			}
			neg = true
			continue
		}
		if set {
			b := buffer.Clear()
			if neg {
				b = -b
			}
			neg = false
			ints = append(ints, b)
			set = false
		}
	}
	if set {
		b := buffer.Clear()
		if neg {
			b = -b
		}
		ints = append(ints, b)
	}
	return ints
}
