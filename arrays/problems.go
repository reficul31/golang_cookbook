package arrays

import (
	"math"
	"sort"
)

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

// Function to find a subarray having maximum sum
// @arg nums - Array of integer to find the subarray
// @returns - Maximum sum of the subarray
func MaxSubArray(nums []int) int {
	maxSum, sum := -int(^uint(0)>>1)-1, 0
	for i := 0; i < len(nums); i = i + 1 {
		sum += nums[i]
		maxSum = max([]int{maxSum, sum})
		sum = max([]int{0, sum})
	}
	return maxSum
}

// Function to find the maximum product in the array
// @arg nums - Array of integers to find the max product
// @returns - Maximum product in the array
func MaxProduct(nums []int) int {
	maxProduct, minProduct := 1, 1
	result := -int(^uint(0)>>1) - 1
	for i := 0; i < len(nums); i = i + 1 {
		tempMax := maxProduct
		maxProduct = max([]int{nums[i], maxProduct * nums[i], minProduct * nums[i]})
		minProduct = min([]int{nums[i], tempMax * nums[i], minProduct * nums[i]})

		result = max([]int{maxProduct, result})
	}
	return result
}

// Function to find the pivot point in the rotated array
// @arg nums - Array of rotated integers
// @returns - Pivot point of the rotation
func FindMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

// Function to find a target element in a rotated array
// @arg nums - Array of rotated integers
// @arg target - Element to be searched in the array
// @returns - Index of the element in the array
func SearchRotatedArray(nums []int, target int) int {
	if len(nums) == 1 && nums[0] == target {
		return 0
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	pivot := left
	left, right = 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		realMid := (mid + pivot) % len(nums)
		if nums[realMid] == target {
			return realMid
		} else if nums[realMid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// Function to find 3 elements with sum equal to 0
// @arg nums - Array of rotated integers
// @returns - Array of indices summing up to 0
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	for i, num := range nums {
		if i != 0 && nums[i-1] == num {
			continue
		}

		left, right := i+1, len(nums)-1
		sum := 0 - num
		for left < right {
			if nums[left]+nums[right] == sum {
				result = append(result, []int{num, nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left += 1
				}
				for left < right && nums[right] == nums[right-1] {
					right -= 1
				}
				left += 1
				right -= 1
			} else if nums[left]+nums[right] < sum {
				left += 1
			} else {
				right -= 1
			}
		}
	}
	return result
}

// Function to find the container with most water
// @arg height - Array of container heights
// @return - Max area of the container in the heights array
func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		if height[left] < height[right] {
			maxArea = max([]int{maxArea, height[left] * (right - left)})
			left += 1
		} else {
			maxArea = max([]int{maxArea, height[right] * (right - left)})
			right -= 1
		}
	}
	return maxArea
}
