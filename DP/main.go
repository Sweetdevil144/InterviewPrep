package main

import (
	"fmt"
)

func main() {
	fmt.Println(numDecodings("1126")) // (1,1,2,6) ; (11, 2, 6) ; (1,1,26) ; (1,12,6) ; (11,26)
}

func numDecodings(s string) int {
	var dfs func(curr string, index int, dp map[int]int) int
	dfs = func(s string, i int, dp map[int]int) int {
		if val, ok := dp[i]; ok {
			return val
		}
		if i == len(s) {
			return 1
		}
		if s[i] == '0' {
			return 0
		}
		res := dfs(s, i+1, dp)
		if i+1 < len(s) && (s[i] == '1' ||
		   (s[i] == '2' && s[i+1] <= '6')) {
			res += dfs(s, i+2, dp)
		}
		dp[i] = res
		return res
	}
	return dfs(s, 0, map[int]int{len(s): 1})
}

func countSubstrings(s string) int {
	l, count := len(s), len(s)
	for i := 0; i < l; i++ {
		a, b := i-1, i+1
		for a > -1 && b < l && s[a] == s[b] {
			count++
			a--
			b++
		}
		a, b = i, i+1
		for a > -1 && b < l && s[a] == s[b] {
			count++
			a--
			b++
		}
	}
	return count
}

func longestPalindrome(s string) string {
	l := len(s)
	res := ""
	for i := 0; i < l; i++ {
		// Odd length palindrome
		a, b := i-1, i+1
		for a > -1 && b < l && s[a] == s[b] {
			a--
			b++
		}
		str := s[a+1 : b]
		if len(str) > len(res) {
			res = str
		}
		// Even length palindrome
		a, b = i, i+1
		for a > -1 && b < l && s[a] == s[b] {
			a--
			b++
		}
		str = s[a+1 : b]
		if len(str) > len(res) {
			res = str
		}
	}
	return res
}

func rob(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	helper := func(nums []int) int {
		l := len(nums)
		dp := make([]int, l+2)
		dp[l+1], dp[l] = 0, 0
		for i := l - 1; i >= 0; i-- {
			dp[i] = max(nums[i]+dp[i+2], dp[i+1])
		}
		return dp[0]
	}
	return max(helper(nums[1:]), helper(nums[:l-1]))
}

func rob1Optimized(nums []int) int {
	l := len(nums)
	dp := make([]int, l+2)
	dp[l+1], dp[l] = 0, 0
	for i := l - 1; i >= 0; i-- {
		dp[i] = max(nums[i]+dp[i+2], dp[i+1])
	}
	return dp[0]
}

func rob1(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	for i := range dp {
		dp[i] = -1
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= l {
			return 0
		}
		if dp[i] != -1 {
			return dp[i]
		}
		dp[i] = max(nums[i]+dfs(i+2), dfs(i+1))
		return dp[i]
	}
	return max(dfs(0), dfs(1))
}

// Input: cost = [1,100,1,1,1,100,1,1,100,1]
// Output: 6
func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	dp := make([]int, l+1)
	for i := range dp {
		dp[i] = -1
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= l {
			return 0
		}
		if dp[i] != -1 {
			return dp[i]
		}
		dp[i] = cost[i] + min(dfs(i+1), dfs(i+2))
		return dp[i]
	}
	return min(dfs(0), dfs(1))
}

func climbStairs(n int) int {
	// n stairs : [1, 2, ]
	ways := []int{1, 2}

	for i := range n {
		ways = append(ways, ways[i]+ways[i+1])
	}
	return ways[n-1]
}

func climbStairsRecursion(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	var backtrack func(i int) int
	backtrack = func(i int) int {
		if i > n {
			return 0
		}
		if i == n {
			return 1
		}
		if dp[i] != 0 {
			return dp[i]
		}
		dp[i] = backtrack(i+1) + backtrack(i+2)
		return dp[i]
	}
	return backtrack(0)
}

func findTargetSumWays(nums []int, target int) int {
	res := 0
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = -1
	}
	var backtrack func(i, currSum int) int
	backtrack = func(i, currSum int) int {
		if i == len(nums) {
			if currSum == target {
				res++
			}
			return res
		}
		if dp[i] != -1 {
			res = dp[i]
		}

		dp[i] = max(backtrack(i+1, currSum+nums[i]), backtrack(i+1, currSum-nums[i]))
		return dp[i]
	}
	res = backtrack(0, 0)
	return res
}
