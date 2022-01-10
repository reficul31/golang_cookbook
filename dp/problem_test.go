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

// Test to find the possible combinations to make a target using nums
func TestCombinationSum4(testCase *testing.T) {
	testCase.Log("To test the possible number of combinations to make a target using nums")

	solution := CombinationSum4([]int{1, 2, 3}, 4)
	if solution != 7 {
		testCase.Errorf("Combination Sum 4 Error: Function returned the wrong value: %d", solution)
	}

	solution = CombinationSum4([]int{9}, 3)
	if solution != 0 {
		testCase.Errorf("Combination Sum 4 Error: Function returned the wrong value: %d", solution)
	}
}

// Test to rob the maximum amount from houses without alerting the police
func TestHouseRobber(testCase *testing.T) {
	testCase.Log("To test the loot of maximum amount from houses without alerting the police")

	solution := HouseRobber([]int{1, 2, 3, 1})
	if solution != 4 {
		testCase.Errorf("House Robber Error: Function returned the wrong value: %d", solution)
	}

	solution = HouseRobber([]int{2, 7, 9, 3, 1})
	if solution != 12 {
		testCase.Errorf("House Robber Error: Function returned the wrong value: %d", solution)
	}
}

// Test to rob the maximum amount from circular houses without alerting the police
func TestHouseRobberCircular(testCase *testing.T) {
	testCase.Log("To test the loot of maximum amount from circular houses without alerting the police")

	solution := HouseRobberCircular([]int{2, 3, 2})
	if solution != 3 {
		testCase.Errorf("House Robber Circular Error: Function returned the wrong value: %d", solution)
	}

	solution = HouseRobberCircular([]int{1, 2, 3, 1})
	if solution != 4 {
		testCase.Errorf("House Robber Circular Error: Function returned the wrong value: %d", solution)
	}

	solution = HouseRobberCircular([]int{1, 2, 3})
	if solution != 3 {
		testCase.Errorf("House Robber Circular Error: Function returned the wrong value: %d", solution)
	}
}

// Test to check the number of decodings possible for a given string of integers
func TestNumberOfDecodings(testCase *testing.T) {
	testCase.Log("To test the number of decodings possible for a given string of integers")

	solution := NumberOfDecodings("12")
	if solution != 2 {
		testCase.Errorf("Number Of Decodings Error: Function returned the wrong value: %d", solution)
	}

	solution = NumberOfDecodings("226")
	if solution != 3 {
		testCase.Errorf("Number Of Decodings Error: Function returned the wrong value: %d", solution)
	}

	solution = NumberOfDecodings("06")
	if solution != 0 {
		testCase.Errorf("Number Of Decodings Error: Function returned the wrong value: %d", solution)
	}
}

// Test to check the total number of ways to reach the bottom right of a matrix
func TestUniquePaths(testCase *testing.T) {
	testCase.Log("To test the total number of ways to reach the bottom right of a matrix")

	solution := UniquePaths(3, 7)
	if solution != 28 {
		testCase.Errorf("Unique Paths Error: Function returned the wrong value: %d", solution)
	}

	solution = UniquePaths(3, 2)
	if solution != 3 {
		testCase.Errorf("Unique Paths Error: Function returned the wrong value: %d", solution)
	}
}

// Test to check whether we can reach the end of the array using the specified number of jumps at the index
func TestCanJump(testCase *testing.T) {
	testCase.Log("To test whether we can reach the end of the array using the specified number of jumps at the index")

	solution := CanJump([]int{3, 2, 1, 0, 4})
	if solution {
		testCase.Errorf("Can Jump Error: Function returned the wrong value: %t", solution)
	}

	solution = CanJump([]int{2, 3, 1, 1, 4})
	if !solution {
		testCase.Errorf("Can Jump Error: Function returned the wrong value: %t", solution)
	}
}
