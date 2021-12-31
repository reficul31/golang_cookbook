package arrays

// Function to get indices in the array that add up to the sum
// @arg nums - Array to search
// @arg target - Target sum to achieve
// @returns - Array of indices
func TwoSum(nums []int, target int) []int {
	hash_map := make(map[int]int)
	for index, num := range nums {
		temp := target - num
		i, ok := hash_map[temp]
		if ok {
			return []int{i, index}
		}

		hash_map[num] = index
	}
	return []int{}
}
