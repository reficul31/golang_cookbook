package sorting

import "testing"

// Test the merge sort function
func TestMergeSort(testCase *testing.T) {
	testCase.Log("To test that the MergeSort function sorts the array correctly")

	arr := []int{2, 3, 5, 4, 1}
	result := MergeSort(arr)

	var prev int
	for index, element := range result {
		if prev > element {
			testCase.Errorf("Sorting Error: Found an element out of order %d at index %d", element, index)
		}
		prev = element
	}
}
