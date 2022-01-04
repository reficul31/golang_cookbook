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

// Test the two sum function
func TestMaxProfit(testCase *testing.T) {
	testCase.Log("To test the max profit from a list of stock prices")

	prices := []int{7, 1, 5, 3, 6, 4}
	solution := MaxProfit(prices)

	if solution != 5 {
		testCase.Errorf("Max Profit Error: Function returned the wrong value")
	}
}

// Test the contains duplicates function
func TestContainsDuplicate(testCase *testing.T) {
	testCase.Log("To test the function to test the duplicates in an array")

	if !ContainsDuplicate([]int{1, 2, 3, 1}) {
		testCase.Errorf("Contains Duplicate Error: Function returned the wrong value")
	}

	if ContainsDuplicate([]int{1, 2, 3, 4}) {
		testCase.Errorf("Contains Duplicate Error: Function returned the wrong value")
	}
}

// Test the product of array except self
func TestProductExceptSelf(testCase *testing.T) {
	testCase.Log("To test the product of array except self")

	nums := []int{1, 2, 3, 4}
	solution := ProductExceptSelf(nums)

	if len(solution) != 4 || solution[0] != 24 || solution[1] != 12 || solution[2] != 8 || solution[3] != 6 {
		testCase.Errorf("Product Except Self Error: Function returned the wrong value")
	}
}

// Test the maximum subarray function
func TestMaxSubArray(testCase *testing.T) {
	testCase.Log("To test the maximum subarray function")

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	solution := MaxSubArray(nums)

	if solution != 6 {
		testCase.Errorf("Max Sub Array Error: Function returned the wrong value")
	}
}

// Test the maximum product of an array
func TestMaxProduct(testCase *testing.T) {
	testCase.Log("To test the maximum product of an array")

	nums := []int{2, 3, -2, 4}
	solution := MaxProduct(nums)

	if solution != 6 {
		testCase.Errorf("Max Sub Array Error: Function returned the wrong value")
	}
}
