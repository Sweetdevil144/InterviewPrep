package main

import (
	"fmt"
	"sort"
)

func firstIndexOf(nums []int, num int) int {
	for i, v := range nums {
		if v == num {
			return i
		}
	}
	return -1
}

func lastIndexOf(nums []int, num int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == num {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(maxFrequency([]int{1, 4, 8, 13}, 5))
}

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	res, n := 1, len(nums)
	if n == 1 {
		return 1
	}
	l, r := 0, 1
	for l <= r && r < n {
		if nums[l] != nums[r] {
			temp, currRes := k, 1
			for i := r - 1; i >= l; i-- {
				if nums[i]-temp+nums[r] >= 0 {
					temp -= nums[r] - nums[i]
					currRes++
				} else {
					break
				}
			}
			res = max(res, currRes)
			l = r
			r++
		} else {
			r++
		}
	}
	return res
}

func maxFrequencyAfterSubarrayOperation(nums []int, k int) (res int) {
	Map := map[int]int{}
	for _, n := range nums {
		Map[n]++
	}
	kadane := func(num int) int {
		var res, currCount int
		for _, n := range nums {
			if n == k {
				currCount--
			}
			if n == num {
				currCount++
			}
			if currCount < 0 {
				currCount = 0
			}
			res = max(res, currCount)
		}
		return res
	}
	for key := range Map {
		res = max(res, kadane(key))
	}
	return res + Map[k]
}

func maxSubArray(nums []int) int {
	right, n := 1, len(nums)
	maxSum, sum := nums[0], nums[0]
	for right < n {
		if sum < 0 {
			sum = nums[right]
		} else {
			sum += nums[right]
		}
		right++
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for _, price := range prices[1:] {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := range nums {
		alt := target - nums[i]
		if _, ok := m[alt]; ok {
			return []int{i, m[alt]}
		}
		m[nums[i]] = i
	}
	return []int{-1, -1}
}
