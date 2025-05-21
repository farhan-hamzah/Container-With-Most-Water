package main

import (
	"fmt"
)

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {
		// Hitung area saat ini
		h := min(height[left], height[right])
		w := right - left
		area := h * w
		if area > maxArea {
			maxArea = area
		}

		// Geser pointer yang lebih kecil
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7})) // Output: 49
}
