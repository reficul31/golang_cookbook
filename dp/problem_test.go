package dp

import "testing"

// Test the total number of distinct ways to climb n steps
func TestClimbStairs(testCase *testing.T) {
	testCase.Log("To test the total number of distinct ways to climb n steps")

	solution := ClimbStairs(3)
	if solution != 3 {
		testCase.Errorf("Climb Stairs Error: Function returned the wrong value: %d", solution)
	}
}

// Test the minimum number of coins required to make the target amount
func TestCoinChange(testCase *testing.T) {
	testCase.Log("To test the minimum number of coins required to make the target amount")

	solution := CoinChange([]int{1, 2, 5}, 11)
	if solution != 3 {
		testCase.Errorf("Coin Change Error: Function returned the wrong value: %d", solution)
	}
}

// Test the length of the longest increasing subsequence
func TestLengthOfLIS(testCase *testing.T) {
	testCase.Log("To test the length of the longest increasing subsequence")

	solution := LengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	if solution != 4 {
		testCase.Errorf("Length of LIS Error: Function returned the wrong value: %d", solution)
	}

	solution = LengthOfLIS([]int{0, 1, 0, 3, 2, 3})
	if solution != 4 {
		testCase.Errorf("Length of LIS Error: Function returned the wrong value: %d", solution)
	}

	solution = LengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7})
	if solution != 1 {
		testCase.Errorf("Length of LIS Error: Function returned the wrong value: %d", solution)
	}
}

// Test the length of the longest common subsequence in two strings
func TestLongestCommonSubsequence(testCase *testing.T) {
	testCase.Log("To test the length of the longest common subsequence in two strings")

	solution := LongestCommonSubsequence("abcde", "ace")
	if solution != 3 {
		testCase.Errorf("Longest Common Subsequence Error: Function returned the wrong value: %d", solution)
	}

	solution = LongestCommonSubsequence("abc", "def")
	if solution != 0 {
		testCase.Errorf("Longest Common Subsequence Error: Function returned the wrong value: %d", solution)
	}
}

// Test whether we can break a word according to the dictionary provided
func TestWordBreak(testCase *testing.T) {
	testCase.Log("To test whether we can break a word according to the dictionary provided")

	solution := WordBreak("leetcode", []string{"leet", "code"})
	if !solution {
		testCase.Errorf("Word Break Error: Function returned the wrong value: %t", solution)
	}

	solution = WordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
	if solution {
		testCase.Errorf("Longest Common Subsequence Error: Function returned the wrong value: %t", solution)
	}
}
