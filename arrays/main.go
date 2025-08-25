package main

import (
	"fmt"
	"math"
	"sort"
	"unicode"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	rows, cols := len(matrix), len(matrix[0])
	firstRowZero := false
	firstColZero := false

	for c := 0; c < cols; c++ {
		if matrix[0][c] == 0 {
			firstRowZero = true
			break
		}
	}
	for r := 0; r < rows; r++ {
		if matrix[r][0] == 0 {
			firstColZero = true
			break
		}
	}

	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if matrix[r][c] == 0 {
				matrix[0][c] = 0
				matrix[r][0] = 0
			}
		}
	}

	for i := 1; i < rows; i++ {
		if matrix[i][0] == 0 {
			for c := 1; c < cols; c++ {
				matrix[i][c] = 0
			}
		}
	}

	for j := 1; j < cols; j++ {
		if matrix[0][j] == 0 {
			for r := 1; r < rows; r++ {
				matrix[r][j] = 0
			}
		}
	}

	if firstRowZero {
		for c := 0; c < cols; c++ {
			matrix[0][c] = 0
		}
	}
	
	if firstColZero {
		for r := 0; r < rows; r++ {
			matrix[r][0] = 0
		}
	}
}

func maxSlidingWindowDP(nums []int, k int) []int {
	n := len(nums)
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = nums[0]
	rightMax[n-1] = nums[n-1]

	for i := 1; i < n; i++ {
		if i%k == 0 {
			leftMax[i] = nums[i]
		} else {
			leftMax[i] = max(leftMax[i-1], nums[i])
		}
	}

	for i := 1; i < n; i++ {
		if (n-1-i)%k == 0 {
			rightMax[n-1-i] = nums[n-1-i]
		} else {
			rightMax[n-1-i] = max(rightMax[n-i], nums[n-1-i])
		}
	}

	output := make([]int, n-k+1)

	for i := 0; i < n-k+1; i++ {
		output[i] = max(leftMax[i+k-1], rightMax[i])
	}
	return output
}

func maxSlidingWindow(nums []int, k int) []int {
	output := []int{}
	q := []int{}
	l, r := 0, 0
	for r < len(nums) {
		for len(q) > 0 && nums[q[len(q)-1]] < nums[r] {
			q = q[:len(q)-1]
		}
		q = append(q, r)
		if l > q[0] {
			q = q[1:]
		}
		if (r + 1) >= k {
			output = append(output, nums[q[0]])
			l += 1
		}
		r += 1
	}
	return output
}

func minWindow(s string, t string) string {
	n, m := len(s), len(t)
	if m > n || n == 0 || m == 0 {
		return ""
	}
	countT := make(map[rune]int)
	for _, ch := range t {
		countT[ch]++
	}
	countS := make(map[rune]int)
	left, right := 0, 0
	required := len(countT)
	match := 0
	minLen := math.MaxInt32
	start, end := 0, 0
	for right < n {
		ch := rune(s[right])
		countS[ch]++
		if count, exists := countT[ch]; exists && countS[ch] == count {
			match++
		}
		for match == required && left <= right {
			currentWindowLen := right - left + 1
			if currentWindowLen < minLen {
				minLen = currentWindowLen
				start = left
				end = right
			}
			leftChar := rune(s[left])
			countS[leftChar]--
			if count, exists := countT[leftChar]; exists && countS[leftChar] < count {
				match--
			}

			left++
		}
		right++
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[start : end+1]
}
func numOfSubarrays(arr []int, k int, threshold int) int {
	res, avg, sum, l := 0, 0, 0, 0
	for i := range k {
		sum += arr[i]
	}
	avg = sum / k
	for i := k; i < len(arr); i++ {
		if avg >= threshold {
			res++
		}
		sum = sum - arr[l] + arr[i]
		l++
		avg = sum / k
	}
	if avg >= threshold {
		res++
	}
	return res
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Count, s2Count, matches, l := make([]int, 26), make([]int, 26), 0, 0
	for i := range len(s1) {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}
	for i := range 26 {
		if s1Count[i] == s2Count[i] {
			matches++
		}
	}
	for r := len(s1); r < len(s2); r++ {
		if matches == 26 {
			return true
		}
		index := s2[r] - 'a'
		s2Count[index]++
		if s1Count[index] == s2Count[index] {
			matches++
		} else if s1Count[index]+1 == s2Count[index] {
			matches--
		}
		index = s2[l] - 'a'
		s2Count[index]--
		if s1Count[index] == s2Count[index] {
			matches++
		} else if s1Count[index]-1 == s2Count[index] {
			matches--
		}
		l++
	}
	return matches == 26
}

func characterReplacement(s string, k int) int {
	count := make([]int, 26)
	res, l, r := 0, 0, 0
	for l <= r && r < len(s) {
		count[byte(s[r])-65]++
		i := getMostFrequentIndex(count)
		if (r - l + 1 - count[i]) <= k {
			res = max(res, r-l+1)
			r++
		} else {
			count[byte(s[l])-65]--
			count[byte(s[r])-65]--
			l++
		}
	}
	return res
}

func getMostFrequentIndex(count []int) int {
	res := 0
	for i, _ := range count {
		if count[i] > count[res] {
			res = i
		}
	}
	return res
}

func maxSubarraySumCircular(nums []int) int {
	currMax := nums[0]
	maxSum := nums[0]

	currMin := nums[0]
	minSum := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		currMax = max(currMax+nums[i], nums[i])
		maxSum = max(maxSum, currMax)

		currMin = min(currMin+nums[i], nums[i])
		minSum = min(minSum, currMin)

		sum += nums[i]
	}

	if sum == minSum {
		return maxSum
	}
	return max(maxSum, sum-minSum)
}

func maxSumSubArray(nums []int) (int, int, int) {
	right, left, n := 1, 0, len(nums)
	maxSum, sum := nums[0], nums[0]
	for right < n {
		if sum < 0 {
			sum = nums[right]
			right++
			left = right
		} else {
			sum += nums[right]
			right++
		}
		maxSum = max(maxSum, sum)
	}
	return maxSum, left, right
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum = max(nums[i], sum+nums[i])
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func lengthOfLongestSubstring(s string) (int, int) {
	start := 0
	maxLen := 0
	seen := make(map[rune]int)
	index := 0

	for _, ch := range s {
		if lastIndex, found := seen[ch]; found && lastIndex >= start {
			start = lastIndex + 1
		}
		seen[ch] = index
		currentLen := index - start + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
		index++
	}
	return maxLen, start
}

func maxProfit_2(prices []int) int {
	profit := 0
	s, e := 0, 1
	for e < len(prices) {
		currentProfit := prices[e] - prices[s]
		if currentProfit > 0 {
			profit += currentProfit
		}
		s = e
		e++
	}
	return profit
}

func maxProfit_1(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			maxProfit = max(maxProfit, prices[i]-minPrice)
		}
	}
	return maxProfit
}

func findMedianSortedArrays(a []int, b []int) float64 {
	if len(a) > len(b) {
		a, b = b, a
	}
	n, m := len(a), len(b)
	total := n + m
	left := (total + 1) / 2

	low, high := 0, n
	for low <= high {
		partitionA := (low + high) / 2
		partitionB := left - partitionA

		var maxLeftA, minRightA, maxLeftB, minRightB int

		if partitionA == 0 {
			maxLeftA = math.MinInt64
		} else {
			maxLeftA = a[partitionA-1]
		}

		if partitionA == n {
			minRightA = math.MaxInt64
		} else {
			minRightA = a[partitionA]
		}

		if partitionB == 0 {
			maxLeftB = math.MinInt64
		} else {
			maxLeftB = b[partitionB-1]
		}

		if partitionB == m {
			minRightB = math.MaxInt64
		} else {
			minRightB = b[partitionB]
		}

		if maxLeftA <= minRightB && maxLeftB <= minRightA {
			if total%2 == 1 {
				return float64(max(maxLeftA, maxLeftB))
			} else {
				return float64(max(maxLeftA, maxLeftB)+min(minRightA, minRightB)) / 2.0
			}
		} else if maxLeftA > minRightB {
			high = partitionA - 1
		} else {
			low = partitionA + 1
		}
	}
	return 0.0
}

func searchTwoPass(nums []int, target int) int {
	point := findMin(nums)
	a1 := binarySearch(nums[:point], target)
	a2 := binarySearch(nums[point:], target)
	if a1 != -1 {
		return a1
	} else if a2 != -1 {
		return a2
	} else {
		return -1
	}
}

func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] {
			if target > nums[mid] || target < nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if target < nums[mid] || target > nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}

func binarySearch(nums []int, target int) int {
	s, e := 0, len(nums)-1
	for s <= e {
		m := s + (e-s)/2
		if nums[m] == target {
			return m
		}
		if nums[m] > target {
			e = m - 1
		} else if nums[m] < target {
			s = m + 1
		}
	}
	return -1
}

func findMin(nums []int) int {
	n := len(nums)
	start, end := 0, n-1
	for start < end {
		mid := (start + end) / 2
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 0
	for _, p := range piles {
		if p > r {
			r = p
		}
	}
	res := r
	for l <= r {
		k := (l + r) / 2
		totalTime := 0
		for _, p := range piles {
			totalTime += int(math.Ceil(float64(p) / float64(k)))
		}
		if totalTime <= h {
			res = k
			r = k - 1
		} else {
			l = k + 1
		}
	}
	return res
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	start, end := 0, m*n-1
	for start <= end {
		mid := start + (end-start)/2
		midValue := matrix[mid/n][mid%n]
		if midValue == target {
			return true
		} else if midValue < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

// O(n) Time, O(1) Memory
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	res, left, right := 0, 0, len(height)-1
	maxLeft, maxRight := 0, 0

	for left <= right {
		if height[left] <= height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right--
		}
	}

	return res
}

// O(N) Memory
func trapFirst(height []int) int {
	res := 0
	mL := maxLeft(height)
	mR := maxRight(height)
	for i := range height {
		res += min(mL[i], mR[i]) - height[i]
	}
	return res
}

func maxLeft(height []int) []int {
	res := make([]int, len(height))
	res[0] = height[0]
	for i := 1; i < len(height); i++ {
		if height[i] > res[i-1] {
			res[i] = height[i]
		} else {
			res[i] = res[i-1]
		}
	}
	return res
}

func maxRight(height []int) []int {
	n := len(height)
	res := make([]int, n)
	res[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		if height[i] > res[i+1] {
			res[i] = height[i]
		} else {
			res[i] = res[i+1]
		}
	}
	return res
}

func maxAreaWater(height []int) int {
	area, i, j := 0, 0, len(height)-1
	for i < j {
		currentArea := (j - i) * min(height[i], height[j])
		area = max(area, currentArea)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return area
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func twoSumSorted(numbers []int, target int) []int {
	s, e := 0, len(numbers)-1
	for s < e {
		if numbers[s]+numbers[e] == target {
			return []int{s + 1, e + 1}
		} else if numbers[s]+numbers[e] < target {
			s++
		} else {
			e--
		}
	}
	return []int{s, e}
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < j && !(unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i]))) {
			i++
		}
		for i < j && !(unicode.IsLetter(rune(s[j])) || unicode.IsDigit(rune(s[j]))) {
			j--
		}
		if i == j {
			break
		}
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
	}
	return true
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	maxStreak := 0
	for num := range set {
		// check if num is start of a subsequence
		if !set[num-1] {
			currentNum := num
			currentStreak := 1
			// continue increase max subsequence length
			for set[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if currentStreak > maxStreak {
				maxStreak = currentStreak
			}
		}
	}
	return maxStreak
}

func isValidSudoku(board [][]byte) bool {
	var rows [9][10]bool
	var cols [9][10]bool
	var boxes [3][3][10]bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			num := board[i][j] - '0'
			if num < 1 || num > 9 {
				return false
			}

			if rows[i][num] || cols[j][num] || boxes[i/3][j/3][num] {
				return false
			}

			rows[i][num] = true
			cols[j][num] = true
			boxes[i/3][j/3][num] = true
		}
	}
	return true
}

func productExceptSelf(nums []int) []int {
	zeroCount := 0
	p1, p2 := 1, 1
	zeroIndex := -1
	res := make([]int, len(nums))

	for i, n := range nums {
		if n == 0 {
			zeroCount++
			zeroIndex = i
			p2 = p1
		} else {
			p1 *= n
			p2 *= n
		}
	}
	if zeroCount > 1 {
		return res
	} else if zeroCount == 1 {
		res[zeroIndex] = p2
		return res
	} else {
		for i, _ := range res {
			res[i] = p1 / nums[i]
		}
		return res
	}
}

func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0)
	m := make(map[int]int)
	arr := make([][]int, 0)

	for _, n := range nums {
		if m[n] > 0 {
			m[n] = m[n] + 1
		} else {
			m[n] = 1
		}
	}
	for key, value := range m {
		arr = append(arr, []int{key, value})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})
	for i := 0; i < k; i++ {
		res = append(res, arr[i][0])
	}
	return res
}

func groupAnagrams(strs []string) [][]string {
	res := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}
		res[count] = append(res[count], s)
	}
	fmt.Println(res)

	var result [][]string
	for _, group := range res {
		result = append(result, group)
	}
	return result
}
