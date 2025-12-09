package util

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Area(p1, p2 [2]int) int {
	w := Abs(p1[0]-p2[0]) + 1
	h := Abs(p1[1]-p2[1]) + 1

	return w * h
}
