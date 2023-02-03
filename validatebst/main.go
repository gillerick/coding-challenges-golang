package main

import (
	"math"
	"testing"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBst(root *TreeNode, minValue, maxValue int) bool {
	if root == nil {
		return true
	}

	if root.Value < minValue || root.Value >= maxValue {
		return false
	}

	isLeftValid := isValidBst(root.Left, minValue, root.Value)
	isRightValid := isValidBst(root.Right, root.Value, maxValue)

	return isRightValid && isLeftValid
}

func TestIsValidBst(t *testing.T) {
	tests := []struct {
		input    *TreeNode
		expected bool
	}{
		{
			input: &TreeNode{
				Value: 15,
				Left: &TreeNode{
					Value: 10,
				},
			},
			expected: true,
		},
		{
			input: &TreeNode{
				Value: 10,
				Left: &TreeNode{
					Value: 17,
				},
				Right: &TreeNode{
					Value: 15,
				},
			},
			expected: false,
		},
		{
			input: &TreeNode{
				Value: 10,
			},
			expected: true,
		},
	}
	for _, test := range tests {
		result := isValidBst(test.input, math.MinInt64, math.MaxInt64)
		if result != test.expected {
			t.Errorf("For tree %v, expected %v but got %v", test.input, test.expected, result)
		}
	}
}
