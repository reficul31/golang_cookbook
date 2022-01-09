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
