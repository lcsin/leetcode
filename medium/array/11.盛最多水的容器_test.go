package array

import (
	"fmt"
	"testing"
)

func TestArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
	fmt.Println(maxAreaByDoublePointer(height))
}

func maxAreaByDoublePointer(height []int) int {
	var h, w, area int
	left, right := 0, len(height)-1
	for i := 0; i < len(height); i++ {
		h = min(height[left], height[right])
		w = right - left
		if h*w > area {
			area = h * w
		}

		switch h {
		case height[left]:
			left++
		case height[right]:
			right--
		}
	}

	return area
}

func maxArea(height []int) int {
	var h, w, area int
	for i := 0; i < len(height); i++ {
		if height[i]*(len(height)-i) < area {
			continue
		}
		for j := i + 1; j < len(height); j++ {
			h = min(height[i], height[j])
			w = j - i
			if h*w > area {
				area = h * w
			}
		}
	}

	return area
}

func min(i, j int) int {
	if i >= j {
		return j
	}
	return i
}
