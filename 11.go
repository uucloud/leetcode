package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func maxArea(height []int) int {
	lp := 0
	rp := len(height) - 1
	m := 0
	for lp != rp {
		t := min(height[rp], height[lp]) * (rp - lp)
		if t > m {
			m = t
		}
		if height[lp] > height[rp] {
			rp--
		} else {
			lp++
		}
	}
	return m
}

func main() {
	print(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
