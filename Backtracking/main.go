package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	res := [][]string{}

	// Helper to create a board from positions
	createBoard := func(queens []int) []string {
		board := make([]string, n)
		for i := 0; i < n; i++ {
			row := make([]byte, n)
			for j := 0; j < n; j++ {
				row[j] = '.'
			}
			row[queens[i]] = 'Q'
			board[i] = string(row)
		}
		return board
	}

	var backtrack func(row int, cols, diag1, diag2 []bool, queens []int)
	backtrack = func(row int, cols, diag1, diag2 []bool, queens []int) {
		if row == n {
			res = append(res, createBoard(queens))
			return
		}
		for col := 0; col < n; col++ {
			d1 := row - col + n - 1
			d2 := row + col
			if cols[col] || diag1[d1] || diag2[d2] {
				continue
			}
			cols[col], diag1[d1], diag2[d2] = true, true, true
			queens = append(queens, col)
			backtrack(row+1, cols, diag1, diag2, queens)
			queens = queens[:len(queens)-1]
			cols[col], diag1[d1], diag2[d2] = false, false, false
		}
	}

	cols := make([]bool, n)
	diag1 := make([]bool, 2*n-1)
	diag2 := make([]bool, 2*n-1)
	backtrack(0, cols, diag1, diag2, []int{})
	return res
}

func letterCombinations(digits string) []string {
	res, l := []string{}, len(digits)
	m := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	var backtrack func(i int, str string)
	if l == 0 {
		return res
	}
	backtrack = func(i int, str string) {
		if i == l {
			res = append(res, str)
			return
		}
		digit := digits[i]
		seq := m[digit]
		for j := range seq {
			str += string(seq[j])
			backtrack(i+1, str)
			str = str[:len(str)-1]
		}
	}
	backtrack(0, "")
	return res
}

func partition(s string) [][]string {
	res, l := make([][]string, 0), len(s)
	var backtrack func(i int, arr []string)
	var isPal = func(str string) bool {
		for i := 0; i < len(str)/2; i++ {
			if str[i] != str[len(str)-i-1] {
				return false
			}
		}
		return true
	}
	backtrack = func(i int, arr []string) {
		if i >= l {
			res = append(res, append([]string{}, arr...))
			return
		}
		for start := i; start < l; start++ {
			temp := s[i : start+1]
			if isPal(temp) {
				arr = append(arr, temp)
				pos := i
				backtrack(start+1, arr)
				arr = arr[:len(arr)-1]
				i = pos
			}
		}
	}
	backtrack(0, []string{})
	return res
}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var backtrack func(i, j, k int) bool
	backtrack = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[k] {
			return false
		}
		tmp := board[i][j]
		board[i][j] = '#'
		found := backtrack(i+1, j, k+1) ||
			backtrack(i-1, j, k+1) ||
			backtrack(i, j+1, k+1) ||
			backtrack(i, j-1, k+1)
		board[i][j] = tmp
		return found
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}
	return false
}

/*
	TODO : Optimize this

Duplicate subset detection: Current use of map[string]bool with string conversion is inefficient;
theoretically, use a more direct representation (e.g., bitmask or index-based) to avoid costly string operations.

Subset generation: Avoid generating and checking all possible subsets for duplicates;
instead, prune duplicates during backtracking by skipping repeated elements (as in combinationSum2).

Space complexity: Reduce extra space by not storing all subsets as strings in a map.

Time complexity: Improve by eliminating redundant work in duplicate checking and subset construction.
*/
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	arr := make([]int, 0)
	sort.Ints(nums)
	m := make(map[string]bool)
	var backtrack func(index int)
	backtrack = func(index int) {
		copy := append([]int{}, arr...)
		str := fmt.Sprintf("%v", copy)
		if !m[str] {
			m[str] = true
			res = append(res, copy)
		}
		for start := index; start < len(nums); start++ {
			arr = append(arr, nums[start])
			backtrack(start + 1)
			arr = arr[:len(arr)-1]
		}
	}
	backtrack(0)
	return res
}

func permute(nums []int) [][]int {
	res := [][]int{}
	var backtrack func(index int, arr []int, m map[int]bool)
	backtrack = func(index int, arr []int, m map[int]bool) {
		if len(arr) == len(nums) {
			res = append(res, append([]int{}, arr...))
			return
		}
		for start := 0; start < len(nums); start++ {
			if !m[nums[start]] {
				arr = append(arr, nums[start])
				m[nums[start]] = true
				backtrack(index+1, arr, m)
				m[nums[start]] = false
				arr = arr[:len(arr)-1]
			}
		}
	}
	backtrack(0, []int{}, make(map[int]bool))
	return res
}

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	var backtrack func(start int, currSum int, arr []int)
	backtrack = func(start int, currSum int, arr []int) {
		if currSum == target {
			res = append(res, append([]int{}, arr...))
			return
		}
		if currSum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			if currSum+candidates[i] > target {
				break
			}
			backtrack(i+1, currSum+candidates[i], append(arr, candidates[i]))
		}
	}
	backtrack(0, 0, []int{})
	return res
}

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	arr := make([]int, 0)
	store := make(map[string]bool)
	var backtrack func(i, currSum int, arr []int)
	backtrack = func(i, currSum int, arr []int) {
		if i == len(candidates) {
			return
		}
		if currSum == target {
			if !store[fmt.Sprintf("%v", arr)] {
				res = append(res, append([]int{}, arr...))
			}
			store[fmt.Sprintf("%v", arr)] = true
			return
		} else if currSum < target {
			backtrack(i, currSum+candidates[i], append(arr, candidates[i]))
			backtrack(i+1, currSum, arr)
		} else {
			return
		}
	}
	backtrack(0, 0, arr)
	return res
}

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	allZero := true
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
		if num != 0 {
			allZero = false
		}
	}
	if allZero {
		return "0"
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})
	res := ""
	for _, s := range strs {
		res += s
	}
	return res
}

func reverse(x int) int {
	var copy int
	isNeg := false
	if x < 0 {
		isNeg = true
		copy = -x
	} else {
		copy = x
	}
	num := 0
	for copy > 0 {
		r := copy % 10
		num = num*10 + r
		copy /= 10
	}
	if isNeg {
		num *= -1
	}
	if num > math.MaxInt32 || num < math.MinInt32-1 {
		return 0
	}
	return num
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	arr := make([]int, 0)
	var backtrack func(index int)
	backtrack = func(index int) {
		res = append(res, append([]int{}, arr...))
		for start := index; start < len(nums); start++ {
			arr = append(arr, nums[start])
			backtrack(start + 1)
			arr = arr[:len(arr)-1]
		}
	}
	backtrack(0)
	return res
}
