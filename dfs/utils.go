package dfs

import "strconv"

func GetLetterFromDigit(digits string, index int) string {
	digit_map := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	num, _ := strconv.Atoi(string(digits[index]))
	letters := digit_map[num]
	return letters
}

func GenerateParenthesisWrapper(n int) []string {
	var result []string
	GenerateParenthesis(n, 0, 0, "", &result)
	return result
}

func LetterCombinationsWrapper(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var result []string
	LetterCombinations(digits, "", &result)
	return result
}

func PermuteWrapper(nums []int) [][]int {
	var result [][]int
	Permutation(nums, []int{}, &result)
	return result
}

func RemoveElement(arr []int, target int) []int {
	new_arr := make([]int, len(arr)-1)
	var index int = 0
	for _, val := range arr {
		if val != target {
			new_arr[index] = val
			index = index + 1
		}
	}
	return new_arr
}

func CombinationSumWrapper(candidates []int, target int) [][]int {
	var result [][]int
	CombinationSum(candidates, 0, target, []int{}, &result)
	return result
}
