package dfs

// Function to find all the letter combinations of a digits using a phone number mapping
func LetterCombinations(digits string, str string, result *[]string) {
	if len(digits) == 0 {
		*result = append(*result, str)
		return
	}

	letters := GetLetterFromDigit(digits, 0)
	for _, letter := range letters {
		str = str + string(letter)
		LetterCombinations(digits[1:], str, result)
		str = str[:len(str)-1]
	}
}

// Function for generating all combinations of the n parenthesis
func GenerateParenthesis(n int, open int, closed int, str string, result *[]string) {
	if len(str) == 2*n {
		*result = append(*result, str)
		return
	}

	if open < n {
		str = str + "("
		GenerateParenthesis(n, open+1, closed, str, result)
		str = str[:len(str)-1]
	}

	if closed < open {
		str = str + ")"
		GenerateParenthesis(n, open, closed+1, str, result)
		str = str[:len(str)-1]
	}
}

// Function to generate all permutations of an array passed to it
func Permutation(nums []int, current []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, current)
		return
	}

	for _, num := range nums {
		current = append(current, num)
		Permutation(RemoveElement(nums, num), current, result)
		current = current[:len(current)-1]
	}
}

// Function to generate all permutation of an array passed to it such that the sum equals the target sum
func CombinationSum(candidates []int, index int, target int, current []int, result *[][]int) {
	if target < 0 {
		return
	}

	if target == 0 {
		*result = append(*result, current)
		return
	}

	for i := index; i < len(candidates); i = i + 1 {
		c := candidates[i]
		if target-c >= 0 {
			current = append(current, c)
			CombinationSum(candidates, i, target-c, current, result)
			current = current[:len(current)-1]
		}
	}
}
