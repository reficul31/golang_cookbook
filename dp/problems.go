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

// Function to check the length of the longest common subsequence in two strings
// @arg text1 - First string (longer string)
// @arg text2 - Second string (shorter string)
// @returns - Length of the longest common subsequence
func LongestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) < len(text2) {
		return LongestCommonSubsequence(text2, text1)
	}

	n, m := len(text1)+1, len(text2)+1
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if text2[i-1] == text1[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max([]int{dp[i-1][j], dp[i][j-1]})
			}
		}
	}
	return dp[m-1][n-1]
}

// Function to find whether we can break a word according to the dictionary provided
// @arg s - String to be broken into component words
// @arg wordDict - Word dictionary
// @return - Boolean that tells whether the word can be broken into the words from the dictionary
func WordBreak(s string, wordDict []string) bool {
	dp, hash_map := make([]bool, len(s)+1), make(map[string]bool)
	for _, word := range wordDict {
		hash_map[word] = true
	}
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] {
				if _, ok := hash_map[s[j:i]]; ok {
					dp[i] = true
				}
			}
		}
	}
	return dp[len(s)]
}

// Function to find the possible combinations to make a target using nums
// @arg nums - The different integers at our disposal
// @arg target - The target sum to be achieved
// @returns - The total numbe of possible combinations to make the target
func CombinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

// Function to find the optimum way to rob houses
// @arg nums - Amount of loot in each house
//  @returns - The maximum amount robbed without alerting the police
func HouseRobber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	two_before, one_before := nums[0], max([]int{nums[0], nums[1]})
	for i := 2; i < len(nums); i++ {
		temp := max([]int{nums[i] + two_before, one_before})
		two_before = one_before
		one_before = temp
	}
	return one_before
}

// Function to find the optimum way to rob circular houses
// @arg nums - Amount of loot in each house
//  @returns - The maximum amount robbed without alerting the police
func HouseRobberCircular(nums []int) int {
	return max([]int{HouseRobber(nums[1:]), HouseRobber(nums[:len(nums)-1])})
}
