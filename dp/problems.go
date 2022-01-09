package dp

// Function to find the total number of ways to reach top
// @arg n - Total steps to reach the top
// @returns - Distinct ways to rach the top
func ClimbStairs(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	one_before, two_before := 2, 1
	total := 0
	for i := 3; i <= n; i++ {
		total = one_before + two_before
		two_before = one_before
		one_before = total
	}
	return total
}
