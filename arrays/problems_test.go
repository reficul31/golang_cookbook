package arrays

import (
	"testing"
)

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

// Test the pivot point of a rotated array
func TestFindMin(testCase *testing.T) {
	testCase.Log("To test the pivot point of a rotated array")

	nums := []int{4, 5, 6, 7, 0, 1, 2}
	solution := FindMin(nums)

	if solution != 0 {
		testCase.Errorf("Find Min Error: Function returned the wrong value")
	}
}

// Test the search in rotated array
func TestSearchRotatedArray(testCase *testing.T) {
	testCase.Log("To test the search in a rotated array")

	nums := []int{4, 5, 6, 7, 0, 1, 2}
	solution := SearchRotatedArray(nums, 5)

	if solution != 1 {
		testCase.Errorf("Search Rotated Array Error: Function returned the wrong value %d", solution)
	}
}

// Test the search in rotated array
func TestThreeSum(testCase *testing.T) {
	testCase.Log("To test function to find 3 elements with sum equal to 0")

	nums := []int{-1, 0, 1, 2, -1, -4}
	solution := ThreeSum(nums)

	if len(solution) != 2 || solution[0][0] != -1 || solution[0][1] != -1 || solution[0][2] != 2 || solution[1][0] != -1 || solution[1][1] != 0 || solution[1][2] != 1 {
		testCase.Errorf("Three Sum Error: Function returned the wrong value")
	}
}

// Test to find the maximum area in the heights array
func TestMaxArea(testCase *testing.T) {
	testCase.Log("To test function to find the maximum area in the heights array")

	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	solution := MaxArea(nums)

	if solution != 49 {
		testCase.Errorf("Max Area Error: Function returned the wrong value %d", solution)
	}
}
