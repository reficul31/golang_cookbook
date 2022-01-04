package arrays

import "math"

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

// Function to find the best time to buy and sell stock
// @arg prices - Array of prices
// @returns - Max profit from the array
func MaxProfit(prices []int) int {
	maxProfit, minPrice := 0, math.MaxInt16
	for _, price := range prices {
		if minPrice > price {
			minPrice = price
		}
		if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}

// Function to find if an array contains duplicates
// @arg nums - Array of integers to find duplicates
func ContainsDuplicate(nums []int) bool {
	hash_map := make(map[int]bool)
	for _, num := range nums {
		if _, ok := hash_map[num]; ok {
			return true
		}
		hash_map[num] = true
	}
	return false
}

// Function to find the product of array exccept self
// @arg nums - Array of integers to find the product
// @returns - Array with the position contains the product except self
func ProductExceptSelf(nums []int) []int {
	product := make([]int, len(nums))
	product[0] = 1
	for i := 1; i < len(nums); i = i + 1 {
		product[i] = product[i-1] * nums[i-1]
	}

	prod := 1
	for i := len(nums) - 1; i >= 0; i = i - 1 {
		product[i] = product[i] * prod
		prod = prod * nums[i]
	}
	return product
}
