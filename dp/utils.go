package dp

func max(arr []int) int {
	max := -int(^uint(0)>>1) - 1
	for _, a := range arr {
		if a > max {
			max = a
		}
	}
	return max
}

func min(arr []int) int {
	min := int(^uint(0) >> 1)
	for _, a := range arr {
		if a < min {
			min = a
		}
	}
	return min
}
