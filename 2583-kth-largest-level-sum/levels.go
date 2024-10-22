package kthlargestlevelsum

import (
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Make tree nodes generic & write a parser for them.
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	var levels []int64
	var traverse func(node *TreeNode, depth int)

	traverse = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(levels) {
			levels = append(levels, 0)
		}
		levels[depth] += int64(node.Val)
		traverse(node.Left, depth+1)
		traverse(node.Right, depth+1)
	}

	traverse(root, 0)
	if len(levels)-k < 0 {
		return -1 // We don't have enough levels
	}
	slices.Sort(levels)
	return levels[len(levels)-k]
}
