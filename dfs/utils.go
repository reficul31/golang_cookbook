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

func SubsetsWrapper(nums []int) [][]int {
	var results [][]int
	Subsets(nums, 0, []int{}, &results)
	return results
}

func Sink(grid *[][]byte, row, column int) {
	if row < 0 || column < 0 || row >= len(*grid) || column >= len((*grid)[row]) || (*grid)[row][column] == '0' {
		return
	}

	(*grid)[row][column] = '0'
	Sink(grid, row+1, column)
	Sink(grid, row-1, column)
	Sink(grid, row, column+1)
	Sink(grid, row, column-1)
}

func CombineWrapper(n int, k int) [][]int {
	var result [][]int
	Combine(n, k, 1, []int{}, &result)
	return result
}
