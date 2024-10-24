package cousinskies

import "solutions/objects"

type TreeNode = objects.TreeNode

// https://leetcode.com/problems/cousins-in-binary-tree-ii/
func replaceValueInTree(root *TreeNode) *TreeNode {
	modifyTree(root, 0, getLevelSums(root))
	return root
}

func modifyTree(root *TreeNode, depth int, levelSums map[int]int) {
	if root == nil {
		return
	}

	modifyTree(root.Left, depth+1, levelSums)
	modifyTree(root.Right, depth+1, levelSums)
	kidsValue := getKidsSum(root) - levelSums[depth]
	setVal(root.Left, kidsValue)
	setVal(root.Right, kidsValue)
}

func getKidsSum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	sum := 0
	if node.Left != nil {
		sum += node.Left.Val
	}
	if node.Right != nil {
		sum += node.Right.Val
	}
	return sum
}

func setVal(node *TreeNode, val int) {
	if node == nil {
		return
	}
	node.Val = val
}

func getLevelSums(root *TreeNode) map[int]int {
	levelSums := map[int]int{}
	var calcLevelSums = func(root *TreeNode, depth int) {}

	calcLevelSums = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		levelSums[depth] += root.Val
		calcLevelSums(root.Left, depth+1)
		calcLevelSums(root.Right, depth+1)
	}
	calcLevelSums(root, 0)

	return levelSums
}
