package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	sum := n * (n + 1) / 2
	var x int
	for i := 1; i < n; i++ {
		fmt.Scan(&x)
		sum -= x
	}
	fmt.Println(sum)
}

func abs(num int) int {
	if num < 0 {
		num *= -1
	}
	return num
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(s, t string) string {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	res := make([]byte, 0, dp[m][n])
	i, j := m, n
	for i > 0 && j > 0 {
		if s[i-1] == t[j-1] {
			res = append(res, s[i-1])
			i--
			j--
		} else if dp[i-1][j] >= dp[i][j-1] {
			i--
		} else {
			j--
		}
	}
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return string(res)
}

// W can be 10^9. memory allocation problems. Need to make dp independent of W & maxW.
func knapsack_2(n, w, maxV int, nums [][]int) int {
	const INF = int(1e18)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, maxV+1)
		for j := 0; j <= maxV; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		a := nums[i-1][0]
		b := nums[i-1][1]
		for j := 0; j <= maxV; j++ {
			dp[i][j] = dp[i-1][j]
			if j-b >= 0 && dp[i-1][j-b] != INF {
				dp[i][j] = min(dp[i][j], dp[i-1][j-b]+a)
			}
		}
	}
	ans := 0
	for v := maxV; v >= 0; v-- {
		if dp[n][v] <= w {
			ans = v
			break
		}
	}
	return ans
}

func knapsack_1(n, w int, nums [][]int) int {
	var backtrack func(index, currW int) int
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, w+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	backtrack = func(index, currW int) int {
		if currW > w || index == n {
			return 0
		}
		if memo[index][currW] > -1 {
			return memo[index][currW]
		}
		if currW+nums[index][0] > w {
			memo[index][currW] = backtrack(index+1, currW)
		} else {
			memo[index][currW] = max(nums[index][1]+backtrack(index+1, currW+nums[index][0]), backtrack(index+1, currW))
		}
		return memo[index][currW]
	}
	return backtrack(0, 0)
}

func getPoints(n int, arr [][]int) int {
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = arr[0][0]
	dp[0][1] = arr[0][1]
	dp[0][2] = arr[0][2]

	for i := 1; i < n; i++ {
		dp[i][0] = arr[i][0] + max(dp[i-1][1], dp[i-1][2])
		dp[i][1] = arr[i][1] + max(dp[i-1][0], dp[i-1][2])
		dp[i][2] = arr[i][2] + max(dp[i-1][0], dp[i-1][1])
	}

	return max(dp[n-1][0], max(dp[n-1][1], dp[n-1][2]))
}

func frogJump(n, k int, arr []int) int {
	const INF = 1 << 60
	dp := make([]int, n)
	for i := range dp {
		dp[i] = INF
	}
	var backtrack func(i int) int
	backtrack = func(i int) int {
		if i == 0 {
			return 0
		}
		if dp[i] != INF {
			return dp[i]
		}
		minCost := INF
		for j := max(0, i-k); j < i; j++ {
			cost := backtrack(j) + abs(arr[j]-arr[i])
			if cost < minCost {
				minCost = cost
			}
		}
		dp[i] = minCost
		return dp[i]
	}
	return backtrack(n - 1)
}

func frogJump_1(n int, arr []int) int {
	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}
	var backtrack func(i int) int
	backtrack = func(i int) int {
		if i == 0 {
			return 0
		}
		if dp[i] != -1 {
			return dp[i]
		}
		cost1 := backtrack(i-1) + abs(arr[i]-arr[i-1])
		cost2 := 1 << 60
		if i > 1 {
			cost2 = backtrack(i-2) + abs(arr[i]-arr[i-2])
		}
		dp[i] = min(cost1, cost2)
		return dp[i]
	}
	return backtrack(n - 1)
}
