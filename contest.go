package main

import "math"

// Area = |x1(y2 - y3) + x2(y3 - y1) + x3(y1 - y2)| / 2
func maxArea(coords [][]int) int64 {
	var maxArea int64 = 0
	n := len(coords)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			base := distance(coords[i], coords[j])

			for k := j + 1; k < n; k++ {
				height := perpendicularDistance(coords[i], coords[j], coords[k])

				currentArea := (base * height)
				maxArea = max(maxArea, currentArea)
			}
		}
	}

	if maxArea == 0 {
		return -1
	}
	return maxArea
}

func distance(p1, p2 []int) int64 {
	dx := p2[0] - p1[0]
	dy := p2[1] - p1[1]
	return int64(math.Sqrt(float64(dx*dx + dy*dy)))
}

func perpendicularDistance(p1, p2, p3 []int) int64 {
	numerator := abs(
		(p2[1]-p1[1])*p3[0] -
			(p2[0]-p1[0])*p3[1] +
			p2[0]*p1[1] -
			p2[1]*p1[0],
	)
	denominator := distance(p1, p2)
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func abs(x int) int64 {
	if x < 0 {
		return int64(-x)
	}
	return int64(x)
}

// TODO
func minSwaps(nums []int) int {
	res := 0
	cE, cO := 0, 0
	for i, _ := range nums {
		if nums[i]%2 == 0 {
			cE++
		} else {
			cO++
		}
	}
	// Base Reject Conditiopns
	if len(nums)%2 == 0 {
		if cE != cO {
			return -1
		}
	} else {
		if math.Abs(float64(cE-cO)) != 1 {
			return -1
		}
	}

	return res
}
