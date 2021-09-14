package sorting

// Function to sort a numerical array using Merge Sort
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := len(arr) / 2
	left_arr := MergeSort(arr[:pivot])
	right_arr := MergeSort(arr[pivot:])
	return Merge(left_arr, right_arr)
}
