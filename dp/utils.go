package dp

func min(arr []int) int {
	min := int(^uint(0) >> 1)
	for _, a := range arr {
		if a < min {
			min = a
		}
	}
	return min
}
