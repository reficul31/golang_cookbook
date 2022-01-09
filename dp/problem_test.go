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
		testCase.Errorf("Climb Stairs Error: Function returned the wrong value: %d", solution)
	}
}
