package dp

import "math"

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

// Function to find the minimum number of coins required to make the amount
// @arg coins - The different denominations given in ascending order
// @arg amount - The target amount to be made using the coins
// @returns - Minimum number of coins that can be used to make the target amount
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
	}

	for i := 0; i <= amount; i++ {
		minChange := math.MaxInt16
		for _, coin := range coins {
			if i-coin < 0 || dp[i-coin] == -1 {
				continue
			}
			minChange = min([]int{dp[i-coin], minChange})
		}

		if minChange != math.MaxInt16 {
			dp[i] = 1 + minChange
		}
	}
	return dp[amount]
}

// Function to find the longest increasing subsequence
// @arg nums - Array to find the longest increasing subsequence
// @returns - Length of the longest increasing subsequence
func LengthOfLIS(nums []int) int {
	dp, maxLength := make([]int, len(nums)), 0
	for i := 1; i < len(nums); i = i + 1 {
		for j := i - 1; j >= 0; j = j - 1 {
			if nums[j] < nums[i] {
				dp[i] = max([]int{dp[i], 1 + dp[j]})
			}
		}
		maxLength = max([]int{maxLength, dp[i]})
	}
	return maxLength + 1
}
