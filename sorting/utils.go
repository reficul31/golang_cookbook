package sorting

// Function to sort two already sorted arrays into a single sorted arrays
func Merge(arr1, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))

	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result[i+j] = arr1[i]
			i = i + 1
		} else {
			result[i+j] = arr2[j]
			j = j + 1
		}
	}

	for i < len(arr1) {
		result[i+j] = arr1[i]
		i = i + 1
	}

	for j < len(arr2) {
		result[i+j] = arr2[j]
		j = j + 1
	}

	return result
}
