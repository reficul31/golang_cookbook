package arrays

import "testing"

// Test the two sum function
func TestTwoSum(testCase *testing.T) {
	testCase.Log("To test the two sum function")

	arr := []int{3, 2, 4}
	target := 6

	solution := TwoSum(arr, target)

	if len(solution) != 2 || solution[0] != 1 || solution[1] != 2 {
		testCase.Errorf("Two Sum Error: Function returned the wrong value")
	}
}
